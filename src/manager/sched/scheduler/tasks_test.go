package scheduler

import (
	"testing"

	"github.com/Dataman-Cloud/swan/src/manager/sched/mock"
	"github.com/Dataman-Cloud/swan/src/mesosproto/mesos"
	"github.com/Dataman-Cloud/swan/src/types"

	"github.com/golang/protobuf/proto"
	"github.com/stretchr/testify/assert"
)

func TestBuildTask(t *testing.T) {
	offer := &mesos.Offer{
		Id: &mesos.OfferID{
			Value: proto.String("abcdefghigklmn"),
		},
		FrameworkId: nil,
		AgentId: &mesos.AgentID{
			Value: proto.String("xxxxxx"),
		},
		Hostname: proto.String("x.x.x.x"),
	}

	version := &types.Version{
		ID:        "test",
		Command:   nil,
		Cpus:      0.1,
		Mem:       16,
		Disk:      0,
		Instances: 5,
		Container: &types.Container{
			Type: "DOCKER",
			Docker: &types.Docker{
				ForcePullImage: proto.Bool(false),
				Image:          proto.String("nginx:1.10"),
				Network:        "BRIDGE",
				Parameters: &[]types.Parameter{
					{
						Key:   "xxxx",
						Value: "yyyy",
					},
				},
				PortMappings: &[]types.PortMapping{
					{
						ContainerPort: 8080,
						Name:          "web",
						Protocol:      "http",
					},
				},
				Privileged: proto.Bool(true),
			},
			Volumes: []*types.Volume{
				{
					ContainerPath: "/tmp/xxxx",
					HostPath:      "/tmp/xxxx",
					Mode:          "RW",
				},
			},
		},
		Labels: &map[string]string{
			"USER_ID": "xxxxx"},
		HealthChecks: []*types.HealthCheck{
			{
				ID:        "xxxxx",
				Address:   "x.x.x.x",
				TaskID:    "aaaa",
				AppID:     "bbbb",
				Port:      nil,
				PortIndex: nil,
				Command:   nil,
				Path:      nil,
			},
		},
		Env: nil,
		KillPolicy: &types.KillPolicy{
			Duration: int64(5),
		},
		UpdatePolicy: nil,
	}

	sched := NewScheduler(FakeConfig(), &mock.Store{})
	task, _ := sched.BuildTask(offer, version, "a.b.c.d")
	assert.Equal(t, task.Name, "a.b.c.d")
}

func TestBuildTaskInfo(t *testing.T) {
	offer := &mesos.Offer{
		Id: &mesos.OfferID{
			Value: proto.String("abcdefghigklmn"),
		},
		FrameworkId: nil,
		AgentId: &mesos.AgentID{
			Value: proto.String("xxxxxx"),
		},
		Hostname: proto.String("x.x.x.x"),
	}

	offer.Resources = append(offer.Resources, &mesos.Resource{
		Name: proto.String("ports"),
		Type: mesos.Value_RANGES.Enum(),
		Ranges: &mesos.Value_Ranges{
			Range: []*mesos.Value_Range{
				{
					Begin: proto.Uint64(uint64(1000)),
					End:   proto.Uint64(uint64(1001)),
				},
			},
		},
	})

	resources := []*mesos.Resource{
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
		&mesos.Resource{
			Name: proto.String("ports"),
			Type: mesos.Value_RANGES.Enum(),
			Ranges: &mesos.Value_Ranges{
				Range: []*mesos.Value_Range{
					{
						Begin: proto.Uint64(uint64(1000)),
						End:   proto.Uint64(uint64(1001)),
					},
				},
			},
		},
	}

	task := &types.Task{
		ID:      "xxxxx",
		Name:    "a.b.c.d",
		Command: nil,
		Cpus:    float64(0.1),
		Mem:     float64(16),
		Disk:    float64(10),
		Image:   proto.String("nginx:1.10"),
		Network: "BRIDGE",
		PortMappings: []*types.PortMappings{
			{
				Port:     8080,
				Protocol: "http",
			},
		},
		Privileged: proto.Bool(false),
		Parameters: []*types.Parameter{
			{
				Key:   "xxxx",
				Value: "yyyy",
			},
		},
		ForcePullImage: proto.Bool(false),
		Volumes: []*types.Volume{
			{
				ContainerPath: "/tmp/xxxx",
				HostPath:      "/tmp/xxxx",
				Mode:          "RW",
			},
		},
		Env:    map[string]string{"DB": "xxxxxxx"},
		Labels: &map[string]string{"USER": "yyyy"},
		HealthChecks: []*types.HealthCheck{
			{
				ID:                  "xxxxx",
				Address:             "x.x.x.x",
				TaskID:              "aaaa",
				AppID:               "bbbb",
				Port:                nil,
				PortIndex:           nil,
				Command:             nil,
				Path:                nil,
				Protocol:            "tcp",
				DelaySeconds:        1,
				ConsecutiveFailures: 1,
				GracePeriodSeconds:  1,
				IntervalSeconds:     1,
				TimeoutSeconds:      1,
			},
		},
		KillPolicy: &types.KillPolicy{
			Duration: int64(5),
		},
		OfferId:       proto.String("xxxyyyzzz"),
		AgentId:       proto.String("mmmnnnzzz"),
		AgentHostname: proto.String("x.x.x.x"),
		Status:        "RUNNING",
		AppId:         "testapp",
	}

	s := NewScheduler(FakeConfig(), &mock.Store{})
	taskInfo := s.BuildTaskInfo(offer, resources, task)
	assert.Equal(t, *taskInfo.Container.Docker.Image, "nginx:1.10")

	task.Network = "NONE"
	taskInfo = s.BuildTaskInfo(offer, resources, task)
	assert.Equal(t, taskInfo.Container.Docker.Network, mesos.ContainerInfo_DockerInfo_NONE.Enum())

	task.Network = "HOST"
	taskInfo = s.BuildTaskInfo(offer, resources, task)
	assert.Equal(t, taskInfo.Container.Docker.Network, mesos.ContainerInfo_DockerInfo_HOST.Enum())
}

func TestLaunchTasks(t *testing.T) {
	offer := &mesos.Offer{
		Id: &mesos.OfferID{
			Value: proto.String("abcdefghigklmn"),
		},
		FrameworkId: nil,
		AgentId: &mesos.AgentID{
			Value: proto.String("xxxxxx"),
		},
		Hostname: proto.String("x.x.x.x"),
	}

	offer.Resources = append(offer.Resources, &mesos.Resource{
		Name: proto.String("ports"),
		Type: mesos.Value_RANGES.Enum(),
		Ranges: &mesos.Value_Ranges{
			Range: []*mesos.Value_Range{
				{
					Begin: proto.Uint64(uint64(1000)),
					End:   proto.Uint64(uint64(1001)),
				},
			},
		},
	})

	resources := []*mesos.Resource{
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
		&mesos.Resource{
			Name: proto.String("ports"),
			Type: mesos.Value_RANGES.Enum(),
			Ranges: &mesos.Value_Ranges{
				Range: []*mesos.Value_Range{
					{
						Begin: proto.Uint64(uint64(1000)),
						End:   proto.Uint64(uint64(1001)),
					},
				},
			},
		},
	}

	task := &types.Task{
		ID:      "xxxxx",
		Name:    "a.b.c.d",
		Command: nil,
		Cpus:    float64(0.1),
		Mem:     float64(16),
		Disk:    float64(10),
		Image:   proto.String("nginx:1.10"),
		Network: "BRIDGE",
		PortMappings: []*types.PortMappings{
			{
				Port:     8080,
				Protocol: "http",
			},
		},
		Privileged: proto.Bool(false),
		Parameters: []*types.Parameter{
			{
				Key:   "xxxx",
				Value: "yyyy",
			},
		},
		ForcePullImage: proto.Bool(false),
		Volumes: []*types.Volume{
			{
				ContainerPath: "/tmp/xxxx",
				HostPath:      "/tmp/xxxx",
				Mode:          "rw",
			},
		},
		Env:    map[string]string{"DB": "xxxxxxx"},
		Labels: &map[string]string{"USER": "yyyy"},
		HealthChecks: []*types.HealthCheck{
			{
				ID:                  "xxxxx",
				Address:             "x.x.x.x",
				TaskID:              "aaaa",
				AppID:               "bbbb",
				Port:                nil,
				PortIndex:           nil,
				Command:             nil,
				Path:                nil,
				Protocol:            "tcp",
				DelaySeconds:        1,
				ConsecutiveFailures: 1,
				GracePeriodSeconds:  1,
				IntervalSeconds:     1,
				TimeoutSeconds:      1,
			},
		},
		KillPolicy: &types.KillPolicy{
			Duration: int64(5),
		},
		OfferId:       proto.String("xxxyyyzzz"),
		AgentId:       proto.String("mmmnnnzzz"),
		AgentHostname: proto.String("x.x.x.x"),
		Status:        "RUNNING",
		AppId:         "testapp",
	}

	s := NewScheduler(FakeConfig(), &mock.Store{})
	taskInfo := s.BuildTaskInfo(offer, resources, task)

	var tasks []*mesos.TaskInfo
	_, err := s.LaunchTasks(offer, append(tasks, taskInfo))
	assert.NotNil(t, err)
}

func TestKillTask(t *testing.T) {
	task := &types.Task{
		ID:      "xxxxx",
		Name:    "a.b.c.d",
		Command: nil,
		Cpus:    float64(0.1),
		Mem:     float64(16),
		Disk:    float64(10),
		Image:   proto.String("nginx:1.10"),
		Network: "BRIDGE",
		PortMappings: []*types.PortMappings{
			{
				Port:     8080,
				Protocol: "http",
			},
		},
		Privileged: proto.Bool(false),
		Parameters: []*types.Parameter{
			{
				Key:   "xxxx",
				Value: "yyyy",
			},
		},
		ForcePullImage: proto.Bool(false),
		Volumes: []*types.Volume{
			{
				ContainerPath: "/tmp/xxxx",
				HostPath:      "/tmp/xxxx",
				Mode:          "rw",
			},
		},
		Env:    map[string]string{"DB": "xxxxxxx"},
		Labels: &map[string]string{"USER": "yyyy"},
		HealthChecks: []*types.HealthCheck{
			{
				ID:        "xxxxx",
				Address:   "x.x.x.x",
				TaskID:    "aaaa",
				AppID:     "bbbb",
				Port:      nil,
				PortIndex: nil,
				Command:   nil,
				Path:      nil,
			},
		},
		KillPolicy: &types.KillPolicy{
			Duration: int64(5),
		},
		OfferId:       proto.String("xxxyyyzzz"),
		AgentId:       proto.String("mmmnnnzzz"),
		AgentHostname: proto.String("x.x.x.x"),
		Status:        "RUNNING",
		AppId:         "testapp",
	}
	s := NewScheduler(FakeConfig(), &mock.Store{})

	_, err := s.KillTask(task)
	assert.NotNil(t, err)
}

func TestReschedulerTask(t *testing.T) {
	s := NewScheduler(FakeConfig(), &mock.Store{})
	go func() {
		s.ReschedulerTask()
	}()

	msg := types.ReschedulerMsg{
		AppID:  "xxxxx",
		TaskID: "yyyyy",
		Err:    make(chan error),
	}

	s.ReschedQueue <- msg
	err := <-msg.Err
	assert.NotNil(t, err)
}
