package main

type ClientState uint8

const (
	ClientStateINIT       ClientState = ClientState(0)
	ClientStateSELECTING  ClientState = ClientState(1)
	ClientStateREQUESTING ClientState = ClientState(2)
	ClientStateBOUND      ClientState = ClientState(3)
	ClientStateRENEWING   ClientState = ClientState(4)
	ClientStateREBINDING  ClientState = ClientState(5)
	ClientStateREBOOTING  ClientState = ClientState(6)
	ClientStateEXPIRED    ClientState = ClientState(7)
	ClientStateDECLINED   ClientState = ClientState(8)
	ClientStateRELEASED   ClientState = ClientState(9)
	ClientStateNIL        ClientState = ClientState(10)
)

var ClientStateToString map[ClientState]string = map[ClientState]string{
	ClientStateINIT:       "INIT",
	ClientStateSELECTING:  "SELECTING",
	ClientStateREQUESTING: "REQUESTING",
	ClientStateBOUND:      "BOUND",
	ClientStateRENEWING:   "RENEWING",
	ClientStateREBINDING:  "REBINDING",
	ClientStateREBOOTING:  "REBOOTING",
	ClientStateEXPIRED:    "EXPIRED",
	ClientStateDECLINED:   "DECLINED",
	ClientStateRELEASED:   "RELEASED",
	ClientStateNIL:        "<nil>",
}

func ClientStateShouldContainNetwork(cs ClientState) bool {
	switch cs {
	case ClientStateBOUND:
		return true
	}

	return false
}

func GetManagementOptionsFromClientState(cs ClientState) {

}
