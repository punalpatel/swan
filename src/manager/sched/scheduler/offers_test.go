package scheduler

import (
	"testing"

	"github.com/Dataman-Cloud/swan/src/manager/sched/mock"
	"github.com/Dataman-Cloud/swan/src/mesosproto/mesos"
	"github.com/Dataman-Cloud/swan/src/mesosproto/sched"

	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
)

func TestRequestOffers(t *testing.T) {
	s := NewScheduler(FakeConfig(), &mock.Store{})

	eventType := sched.Event_OFFERS

	offer := mesos.Offer{
		Id: &mesos.OfferID{
			Value: proto.String("abcdefghigklmn"),
		},
		FrameworkId: nil,
		AgentId: &mesos.AgentID{
			Value: proto.String("xxxxxx"),
		},
		Hostname: proto.String("x.x.x.x"),
	}
	event := &sched.Event{
		Offers: &sched.Event_Offers{
			Offers: []*mesos.Offer{
				&offer,
			},
		},
	}

	s.AddEvent(eventType, event)

	offers, _ := s.RequestOffers(nil)
	assert.Equal(t, *offers[0].Id.Value, "abcdefghigklmn")

	s.AddEvent(eventType, nil)
	offers, _ = s.RequestOffers(nil)
	assert.Nil(t, offers)
}

func TestOfferedResources(t *testing.T) {
	offer := mesos.Offer{
		Resources: []*mesos.Resource{
			&mesos.Resource{
				Name:   proto.String("cpus"),
				Type:   mesos.Value_SCALAR.Enum(),
				Scalar: &mesos.Value_Scalar{Value: proto.Float64(0.1)},
			},
			&mesos.Resource{
				Name:   proto.String("mem"),
				Type:   mesos.Value_SCALAR.Enum(),
				Scalar: &mesos.Value_Scalar{Value: proto.Float64(16)},
			},
			&mesos.Resource{
				Name:   proto.String("disk"),
				Type:   mesos.Value_SCALAR.Enum(),
				Scalar: &mesos.Value_Scalar{Value: proto.Float64(10)},
			},
		},
	}

	s := NewScheduler(FakeConfig(), &mock.Store{})
	cpus, mem, disk := s.OfferedResources(&offer)

	assert.Equal(t, cpus, float64(0.1))
	assert.Equal(t, mem, float64(16))
	assert.Equal(t, disk, float64(10))
}

func TestDeclineResource(t *testing.T) {
	s := NewScheduler(FakeConfig(), &mock.Store{})
	_, err := s.DeclineResource(proto.String("xxxxx-yyyyy-zzzzz"))
	assert.NotNil(t, err)
}
