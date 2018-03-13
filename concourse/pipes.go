package concourse

import (
	"github.com/concourse/atc"
	"github.com/concourse/go-concourse/concourse/internal"
	"github.com/tedsuo/rata"
)

func (team *team) CreatePipe() (atc.Pipe, error) {
	var pipe atc.Pipe
	err := team.connection.Send(internal.Request{
		RequestName: atc.CreatePipe,
		Params:      rata.Params{"team_name": team.Name()},
	}, &internal.Response{
		Result: &pipe,
	})

	return pipe, err
}
