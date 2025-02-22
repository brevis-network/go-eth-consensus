package consensus

type SignedBeaconBlock interface {
	isSignedBeaconBlock()
}

func (s *SignedBeaconBlockAltair) isSignedBeaconBlock() {
}

func (s *SignedBeaconBlockPhase0) isSignedBeaconBlock() {
}

func (s *SignedBeaconBlockBellatrix) isSignedBeaconBlock() {
}

func (s *SignedBeaconBlockCapella) isSignedBeaconBlock() {
}

func (s *SignedBeaconBlockDeneb) isSignedBeaconBlock() {
}

type BeaconBlock interface {
	isBeaconBlock()
}

func (s *BeaconBlockAltair) isBeaconBlock() {
}

func (s *BeaconBlockPhase0) isBeaconBlock() {
}

func (s *BeaconBlockBellatrix) isBeaconBlock() {
}

func (s *BeaconBlockCapella) isBeaconBlock() {
}

func (s *BeaconBlockDeneb) isBeaconBlock() {
}

type BeaconState interface {
	isBeaconState()
}

func (s *BeaconStatePhase0) isBeaconState() {
}

func (s *BeaconStateAltair) isBeaconState() {
}

func (s *BeaconStateBellatrix) isBeaconState() {
}

func (s *BeaconStateCapella) isBeaconState() {
}

func (s *BeaconStateDeneb) isBeaconState() {
}
