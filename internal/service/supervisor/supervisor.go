package supervisor

import (
	"encoding/json"
	"github.com/foxfurry/go_delivery/internal/domain/dto"
	"github.com/foxfurry/go_delivery/internal/infrastracture/logger"
	"github.com/foxfurry/go_delivery/internal/service/supervisor/gateway"
)

type ISupervisor interface {
	Order(order *dto.ClientOrder) error
	Menu() (*dto.Menu, error)
	Distribution(clientID dto.Distribution) error
}

type deliverySupervisor struct {
	gate gateway.IGateway
	clients map[int]int					// Map from client ID to client order count
}

func NewDeliverySupervisor() ISupervisor {
	return &deliverySupervisor{
		gate: gateway.NewDeliveryGateway(),
		clients: make(map[int]int),
	}
}

func (s *deliverySupervisor) Distribution(distribute dto.Distribution) error {
	logger.LogSuperF("Received distribution. Client %d has %d suborders remaining", distribute.ClientID, s.clients[distribute.ClientID]-1)

	s.clients[distribute.ClientID]--

	if s.clients[distribute.ClientID] == 0 {
		logger.LogSuperF("All suborders complete. Distributing order to client %d", distribute.ClientID)
		s.gate.Distribute(distribute)
	}

	return nil
}

func (s *deliverySupervisor) Order(order *dto.ClientOrder) error {
	logger.LogSuperF("Received order from %d with %d suborders", order.ClientID, len(order.Orders))
	s.clients[order.ClientID] = len(order.Orders)

	for _, val := range order.Orders {
		val.ClientID = order.ClientID
		s.gate.SendOrder(val)
	}

	return nil
}

func (s *deliverySupervisor) Menu() (*dto.Menu, error) {
	resp, err := s.gate.GetMenu()
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	menuHolder := new(dto.Menu)

	err = json.NewDecoder(resp.Body).Decode(menuHolder)
	if err != nil {
		return nil, err
	}

	return menuHolder, err
}
