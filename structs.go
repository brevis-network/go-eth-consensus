package consensus

type Root [32]byte

type Signature [96]byte

type AggregateAndProof struct {
	Index          uint64       `json:"aggregator_index"`
	Aggregate      *Attestation `json:"aggregate"`
	SelectionProof [96]byte     `json:"selection_proof" ssz-size:"96"`
}

type Checkpoint struct {
	Epoch uint64 `json:"epoch"`
	Root  Root   `json:"root" ssz-size:"32"`
}

type AttestationData struct {
	Slot            uint64      `json:"slot"`
	Index           uint64      `json:"index"`
	BeaconBlockHash [32]byte    `json:"beacon_block_root" ssz-size:"32"`
	Source          *Checkpoint `json:"source"`
	Target          *Checkpoint `json:"target"`
}

type Attestation struct {
	AggregationBits []byte           `json:"aggregation_bits" ssz:"bitlist" ssz-max:"2048"`
	Data            *AttestationData `json:"data"`
	Signature       Signature        `json:"signature" ssz-size:"96"`
}

type DepositData struct {
	Pubkey                [48]byte  `json:"pubkey" ssz-size:"48"`
	WithdrawalCredentials [32]byte  `json:"withdrawal_credentials" ssz-size:"32"`
	Amount                uint64    `json:"amount"`
	Signature             Signature `json:"signature" ssz-size:"96"`
	Root                  [32]byte  `ssz:"-"`
}

type Deposit struct {
	Proof [33][32]byte `ssz-size:"33,32"`
	Data  *DepositData
}

type DepositMessage struct {
	Pubkey                [48]byte `json:"pubkey" ssz-size:"48"`
	WithdrawalCredentials [32]byte `json:"withdrawal_credentials" ssz-size:"32"`
	Amount                uint64   `json:"amount"`
}

type IndexedAttestation struct {
	AttestationIndices []uint64         `json:"attesting_indices" ssz-max:"2048"`
	Data               *AttestationData `json:"data"`
	Signature          Signature        `json:"signature" ssz-size:"96"`
}

type PendingAttestation struct {
	AggregationBits []byte           `json:"aggregation_bits" ssz:"bitlist" ssz-max:"2048"`
	Data            *AttestationData `json:"data"`
	InclusionDelay  uint64           `json:"inclusion_delay"`
	ProposerIndex   uint64           `json:"proposer_index"`
}

type Fork struct {
	PreviousVersion [4]byte `json:"previous_version" ssz-size:"4"`
	CurrentVersion  [4]byte `json:"current_version" ssz-size:"4"`
	Epoch           uint64  `json:"epoch"`
}

type Validator struct {
	Pubkey                     [48]byte `json:"pubkey" ssz-size:"48"`
	WithdrawalCredentials      [32]byte `json:"withdrawal_credentials" ssz-size:"32"`
	EffectiveBalance           uint64   `json:"effective_balance"`
	Slashed                    bool     `json:"slashed"`
	ActivationEligibilityEpoch uint64   `json:"activation_eligibility_epoch"`
	ActivationEpoch            uint64   `json:"activation_epoch"`
	ExitEpoch                  uint64   `json:"exit_epoch"`
	WithdrawableEpoch          uint64   `json:"withdrawable_epoch"`
}

type VoluntaryExit struct {
	Epoch          uint64 `json:"epoch"`
	ValidatorIndex uint64 `json:"validator_index"`
}

type SignedVoluntaryExit struct {
	Exit      *VoluntaryExit `json:"message"`
	Signature Signature      `json:"signature" ssz-size:"96"`
}

type HistoricalBatch struct {
	BlockRoots [8192][32]byte `json:"block_roots" ssz-size:"8192,32"`
	StateRoots [8192][32]byte `json:"state_roots" ssz-size:"8192,32"`
}

type Eth1Data struct {
	DepositRoot  Root     `json:"deposit_root" ssz-size:"32"`
	DepositCount uint64   `json:"deposit_count"`
	BlockHash    [32]byte `json:"block_hash" ssz-size:"32"`
}

type SigningRoot struct {
	ObjectRoot Root   `json:"object_root" ssz-size:"32"`
	Domain     []byte `json:"domain" ssz-size:"8"`
}

type ProposerSlashing struct {
	Header1 *SignedBeaconBlockHeader `json:"signed_header_1"`
	Header2 *SignedBeaconBlockHeader `json:"signed_header_2"`
}

type AttesterSlashing struct {
	Attestation1 *IndexedAttestation `json:"attestation_1"`
	Attestation2 *IndexedAttestation `json:"attestation_2"`
}

type Transfer struct {
	Sender    uint64    `json:"sender"`
	Recipient uint64    `json:"recipient"`
	Amount    uint64    `json:"amount"`
	Fee       uint64    `json:"fee"`
	Slot      uint64    `json:"slot"`
	Pubkey    [48]byte  `json:"pubkey" ssz-size:"48"`
	Signature Signature `json:"signature" ssz-size:"96"`
}

type BeaconStatePhase0 struct {
	GenesisTime                 uint64                `json:"genesis_time"`
	GenesisValidatorsRoot       [32]byte              `json:"genesis_validators_root" ssz-size:"32"`
	Slot                        uint64                `json:"slot"`
	Fork                        *Fork                 `json:"fork"`
	LatestBlockHeader           *BeaconBlockHeader    `json:"latest_block_header"`
	BlockRoots                  [8192][32]byte        `json:"block_roots" ssz-size:"8192,32"`
	StateRoots                  [8192][32]byte        `json:"state_roots" ssz-size:"8192,32"`
	HistoricalRoots             [][32]byte            `json:"historical_roots" ssz-max:"16777216" ssz-size:"?,32"`
	Eth1Data                    *Eth1Data             `json:"eth1_data"`
	Eth1DataVotes               []*Eth1Data           `json:"eth1_data_votes" ssz-max:"2048"`
	Eth1DepositIndex            uint64                `json:"eth1_deposit_index"`
	Validators                  []*Validator          `json:"validators" ssz-max:"1099511627776"`
	Balances                    []uint64              `json:"balances" ssz-max:"1099511627776"`
	RandaoMixes                 [65536][32]byte       `json:"randao_mixes" ssz-size:"65536,32"`
	Slashings                   []uint64              `json:"slashings" ssz-size:"8192"`
	PreviousEpochAttestations   []*PendingAttestation `json:"previous_epoch_attestations" ssz-max:"4096"`
	CurrentEpochAttestations    []*PendingAttestation `json:"current_epoch_attestations" ssz-max:"4096"`
	JustificationBits           [1]byte               `json:"justification_bits" ssz-size:"1"`
	PreviousJustifiedCheckpoint *Checkpoint           `json:"previous_justified_checkpoint"`
	CurrentJustifiedCheckpoint  *Checkpoint           `json:"current_justified_checkpoint"`
	FinalizedCheckpoint         *Checkpoint           `json:"finalized_checkpoint"`
}

type SignedBeaconBlockPhase0 struct {
	Block     *BeaconBlockPhase0 `json:"message"`
	Signature Signature          `json:"signature" ssz-size:"96"`
}

type BeaconBlockPhase0 struct {
	Slot          uint64                 `json:"slot"`
	ProposerIndex uint64                 `json:"proposer_index"`
	ParentRoot    Root                   `json:"parent_root" ssz-size:"32"`
	StateRoot     Root                   `json:"state_root" ssz-size:"32"`
	Body          *BeaconBlockBodyPhase0 `json:"body"`
}

type BeaconBlockBodyPhase0 struct {
	RandaoReveal      Signature              `json:"randao_reveal" ssz-size:"96"`
	Eth1Data          *Eth1Data              `json:"eth1_data"`
	Graffiti          [32]byte               `json:"graffiti" ssz-size:"32"`
	ProposerSlashings []*ProposerSlashing    `json:"proposer_slashings" ssz-max:"16"`
	AttesterSlashings []*AttesterSlashing    `json:"attester_slashings" ssz-max:"2"`
	Attestations      []*Attestation         `json:"attestations" ssz-max:"128"`
	Deposits          []*Deposit             `json:"deposits" ssz-max:"16"`
	VoluntaryExits    []*SignedVoluntaryExit `json:"voluntary_exits" ssz-max:"16"`
}

type SignedBeaconBlockHeader struct {
	Header    *BeaconBlockHeader `json:"message"`
	Signature Signature          `json:"signature" ssz-size:"96"`
}

type BeaconBlockHeader struct {
	Slot          uint64 `json:"slot"`
	ProposerIndex uint64 `json:"proposer_index"`
	ParentRoot    Root   `json:"parent_root" ssz-size:"32"`
	StateRoot     Root   `json:"state_root" ssz-size:"32"`
	BodyRoot      Root   `json:"body_root" ssz-size:"32"`
}

type ForkData struct {
	CurrentVersion        [4]byte `json:"current_version" ssz-size:"4"`
	GenesisValidatorsRoot Root    `json:"genesis_validators_root" ssz-size:"32"`
}

type SigningData struct {
	ObjectRoot Root     `json:"object_root" ssz-size:"32"`
	Domain     [32]byte `json:"domain" ssz-size:"32"`
}

// Altair fork

type LightClientHeader struct {
	Header *BeaconBlockHeader `json:"beacon"`
}

type LightClientBootstrap struct {
	Header                     *LightClientHeader `json:"header"`
	CurrentSyncCommittee       *SyncCommittee     `json:"current_sync_committee"`
	CurrentSyncCommitteeBranch [][32]byte         `json:"current_sync_committee_branch" ssz-size:"5,32"`
}

type LightClientFinalityUpdate struct {
	AttestedHeader  *LightClientHeader `json:"attested_header"`
	FinalizedHeader *LightClientHeader `json:"finalized_header"`
	FinalityBranch  [][32]byte         `json:"finality_branch" ssz-size:"6,32"`
	SyncAggregate   *SyncAggregate     `json:"sync_aggregate"`
	SignatureSlot   uint64             `json:"signature_slot"`
}

type LightClientOptimisticUpdate struct {
	AttestedHeader *LightClientHeader `json:"attested_header"`
	SyncAggregate  *SyncAggregate     `json:"sync_aggregate"`
	SignatureSlot  uint64             `json:"signature_slot"`
}

type LightClientUpdate struct {
	AttestedHeader          *LightClientHeader `json:"attested_header"`
	NextSyncCommittee       *SyncCommittee     `json:"next_sync_committee"`
	NextSyncCommitteeBranch [][32]byte         `json:"next_sync_committee_branch" ssz-size:"5,32"`
	FinalizedHeader         *LightClientHeader `json:"finalized_header"`
	FinalityBranch          [][32]byte         `json:"finality_branch" ssz-size:"6,32"`
	SyncAggregate           *SyncAggregate     `json:"sync_aggregate"`
	SignatureSlot           uint64             `json:"signature_slot"`
}

type BeaconStateAltair struct {
	GenesisTime                 uint64             `json:"genesis_time"`
	GenesisValidatorsRoot       [32]byte           `json:"genesis_validators_root" ssz-size:"32"`
	Slot                        uint64             `json:"slot"`
	Fork                        *Fork              `json:"fork"`
	LatestBlockHeader           *BeaconBlockHeader `json:"latest_block_header"`
	BlockRoots                  [8192][32]byte     `json:"block_roots" ssz-size:"8192,32"`
	StateRoots                  [8192][32]byte     `json:"state_roots" ssz-size:"8192,32"`
	HistoricalRoots             [][32]byte         `json:"historical_roots" ssz-max:"16777216" ssz-size:"?,32"`
	Eth1Data                    *Eth1Data          `json:"eth1_data"`
	Eth1DataVotes               []*Eth1Data        `json:"eth1_data_votes" ssz-max:"2048"`
	Eth1DepositIndex            uint64             `json:"eth1_deposit_index"`
	Validators                  []*Validator       `json:"validators" ssz-max:"1099511627776"`
	Balances                    []uint64           `json:"balances" ssz-max:"1099511627776"`
	RandaoMixes                 [65536][32]byte    `json:"randao_mixes" ssz-size:"65536,32"`
	Slashings                   []uint64           `json:"slashings" ssz-size:"8192"`
	PreviousEpochParticipation  []byte             `json:"previous_epoch_participation" ssz-max:"1099511627776"`
	CurrentEpochParticipation   []byte             `json:"current_epoch_participation" ssz-max:"1099511627776"`
	JustificationBits           [1]byte            `json:"justification_bits" ssz-size:"1"`
	PreviousJustifiedCheckpoint *Checkpoint        `json:"previous_justified_checkpoint"`
	CurrentJustifiedCheckpoint  *Checkpoint        `json:"current_justified_checkpoint"`
	FinalizedCheckpoint         *Checkpoint        `json:"finalized_checkpoint"`
	InactivityScores            []uint64           `json:"inactivity_scores" ssz-max:"1099511627776"`
	CurrentSyncCommittee        *SyncCommittee     `json:"current_sync_committee"`
	NextSyncCommittee           *SyncCommittee     `json:"next_sync_committee"`
}

type SignedBeaconBlockAltair struct {
	Block     *BeaconBlockAltair `json:"message"`
	Signature Signature          `json:"signature" ssz-size:"96"`
}

type BeaconBlockAltair struct {
	Slot          uint64                 `json:"slot"`
	ProposerIndex uint64                 `json:"proposer_index"`
	ParentRoot    Root                   `json:"parent_root" ssz-size:"32"`
	StateRoot     Root                   `json:"state_root" ssz-size:"32"`
	Body          *BeaconBlockBodyAltair `json:"body"`
}

type BeaconBlockBodyAltair struct {
	RandaoReveal      Signature              `json:"randao_reveal" ssz-size:"96"`
	Eth1Data          *Eth1Data              `json:"eth1_data"`
	Graffiti          [32]byte               `json:"graffiti" ssz-size:"32"`
	ProposerSlashings []*ProposerSlashing    `json:"proposer_slashings" ssz-max:"16"`
	AttesterSlashings []*AttesterSlashing    `json:"attester_slashings" ssz-max:"2"`
	Attestations      []*Attestation         `json:"attestations" ssz-max:"128"`
	Deposits          []*Deposit             `json:"deposits" ssz-max:"16"`
	VoluntaryExits    []*SignedVoluntaryExit `json:"voluntary_exits" ssz-max:"16"`
	SyncAggregate     *SyncAggregate         `json:"sync_aggregate"`
}

type SyncAggregate struct {
	SyncCommiteeBits      [64]byte  `json:"sync_committee_bits" ssz-size:"64"`
	SyncCommiteeSignature Signature `json:"sync_committee_signature" ssz-size:"96"`
}

type SyncCommittee struct {
	PubKeys         [512][48]byte `json:"pubkeys" ssz-size:"512,48"`
	AggregatePubKey [48]byte      `json:"aggregate_pubkey" ssz-size:"48"`
}

// bellatrix

type BeaconStateBellatrix struct {
	GenesisTime                  uint64                  `json:"genesis_time"`
	GenesisValidatorsRoot        [32]byte                `json:"genesis_validators_root" ssz-size:"32"`
	Slot                         uint64                  `json:"slot"`
	Fork                         *Fork                   `json:"fork"`
	LatestBlockHeader            *BeaconBlockHeader      `json:"latest_block_header"`
	BlockRoots                   [8192][32]byte          `json:"block_roots" ssz-size:"8192,32"`
	StateRoots                   [8192][32]byte          `json:"state_roots" ssz-size:"8192,32"`
	HistoricalRoots              [][]byte                `json:"historical_roots" ssz-max:"16777216" ssz-size:"?,32"`
	Eth1Data                     *Eth1Data               `json:"eth1_data"`
	Eth1DataVotes                []*Eth1Data             `json:"eth1_data_votes" ssz-max:"2048"`
	Eth1DepositIndex             uint64                  `json:"eth1_deposit_index"`
	Validators                   []*Validator            `json:"validators" ssz-max:"1099511627776"`
	Balances                     []uint64                `json:"balances" ssz-max:"1099511627776"`
	RandaoMixes                  [65536][32]byte         `json:"randao_mixes" ssz-size:"65536,32"`
	Slashings                    []uint64                `json:"slashings" ssz-size:"8192"`
	PreviousEpochParticipation   []byte                  `json:"previous_epoch_participation" ssz-max:"1099511627776"`
	CurrentEpochParticipation    []byte                  `json:"current_epoch_participation" ssz-max:"1099511627776"`
	JustificationBits            [1]byte                 `json:"justification_bits" ssz-size:"1"`
	PreviousJustifiedCheckpoint  *Checkpoint             `json:"previous_justified_checkpoint"`
	CurrentJustifiedCheckpoint   *Checkpoint             `json:"current_justified_checkpoint"`
	FinalizedCheckpoint          *Checkpoint             `json:"finalized_checkpoint"`
	InactivityScores             []uint64                `json:"inactivity_scores" ssz-max:"1099511627776"`
	CurrentSyncCommittee         *SyncCommittee          `json:"current_sync_committee"`
	NextSyncCommittee            *SyncCommittee          `json:"next_sync_committee"`
	LatestExecutionPayloadHeader *ExecutionPayloadHeader `json:"latest_execution_payload_header"`
}

type SignedBeaconBlockBellatrix struct {
	Block     *BeaconBlockBellatrix `json:"message"`
	Signature Signature             `json:"signature" ssz-size:"96"`
}

type BeaconBlockBellatrix struct {
	Slot          uint64                    `json:"slot"`
	ProposerIndex uint64                    `json:"proposer_index"`
	ParentRoot    Root                      `json:"parent_root" ssz-size:"32"`
	StateRoot     Root                      `json:"state_root" ssz-size:"32"`
	Body          *BeaconBlockBodyBellatrix `json:"body"`
}

type BeaconBlockBodyBellatrix struct {
	RandaoReveal      Signature              `json:"randao_reveal" ssz-size:"96"`
	Eth1Data          *Eth1Data              `json:"eth1_data"`
	Graffiti          [32]byte               `json:"graffiti" ssz-size:"32"`
	ProposerSlashings []*ProposerSlashing    `json:"proposer_slashings" ssz-max:"16"`
	AttesterSlashings []*AttesterSlashing    `json:"attester_slashings" ssz-max:"2"`
	Attestations      []*Attestation         `json:"attestations" ssz-max:"128"`
	Deposits          []*Deposit             `json:"deposits" ssz-max:"16"`
	VoluntaryExits    []*SignedVoluntaryExit `json:"voluntary_exits" ssz-max:"16"`
	SyncAggregate     *SyncAggregate         `json:"sync_aggregate"`
	ExecutionPayload  *ExecutionPayload      `json:"execution_payload"`
}

type SignedBlindedBeaconBlock struct {
	Block     *BlindedBeaconBlock `json:"message"`
	Signature Signature           `json:"signature" ssz-size:"96"`
}

type BlindedBeaconBlock struct {
	Slot          uint64                  `json:"slot"`
	ProposerIndex uint64                  `json:"proposer_index"`
	ParentRoot    Root                    `json:"parent_root" ssz-size:"32"`
	StateRoot     Root                    `json:"state_root" ssz-size:"32"`
	Body          *BlindedBeaconBlockBody `json:"body"`
}

type BlindedBeaconBlockBody struct {
	RandaoReveal           Signature               `json:"randao_reveal" ssz-size:"96"`
	Eth1Data               *Eth1Data               `json:"eth1_data"`
	Graffiti               [32]byte                `json:"graffiti" ssz-size:"32"`
	ProposerSlashings      []*ProposerSlashing     `json:"proposer_slashings" ssz-max:"16"`
	AttesterSlashings      []*AttesterSlashing     `json:"attester_slashings" ssz-max:"2"`
	Attestations           []*Attestation          `json:"attestations" ssz-max:"128"`
	Deposits               []*Deposit              `json:"deposits" ssz-max:"16"`
	VoluntaryExits         []*SignedVoluntaryExit  `json:"voluntary_exits" ssz-max:"16"`
	SyncAggregate          *SyncAggregate          `json:"sync_aggregate"`
	ExecutionPayloadHeader *ExecutionPayloadHeader `json:"execution_payload_header"`
}

type ExecutionPayload struct {
	ParentHash    [32]byte  `ssz-size:"32" json:"parent_hash"`
	FeeRecipient  [20]byte  `ssz-size:"20" json:"fee_recipient"`
	StateRoot     [32]byte  `ssz-size:"32" json:"state_root"`
	ReceiptsRoot  [32]byte  `ssz-size:"32" json:"receipts_root"`
	LogsBloom     [256]byte `ssz-size:"256" json:"logs_bloom"`
	PrevRandao    [32]byte  `ssz-size:"32" json:"prev_randao"`
	BlockNumber   uint64    `json:"block_number"`
	GasLimit      uint64    `json:"gas_limit"`
	GasUsed       uint64    `json:"gas_used"`
	Timestamp     uint64    `json:"timestamp"`
	ExtraData     []byte    `ssz-max:"32" json:"extra_data"`
	BaseFeePerGas Uint256   `ssz-size:"32" json:"base_fee_per_gas"`
	BlockHash     [32]byte  `ssz-size:"32" json:"block_hash"`
	Transactions  [][]byte  `ssz-max:"1048576,1073741824" ssz-size:"?,?" json:"transactions"`
}

type Uint256 [32]byte

type ExecutionPayloadHeader struct {
	ParentHash       [32]byte  `json:"parent_hash" ssz-size:"32"`
	FeeRecipient     [20]byte  `json:"fee_recipient" ssz-size:"20"`
	StateRoot        [32]byte  `json:"state_root" ssz-size:"32"`
	ReceiptsRoot     [32]byte  `json:"receipts_root" ssz-size:"32"`
	LogsBloom        [256]byte `json:"logs_bloom" ssz-size:"256"`
	PrevRandao       [32]byte  `json:"prev_randao" ssz-size:"32"`
	BlockNumber      uint64    `json:"block_number"`
	GasLimit         uint64    `json:"gas_limit"`
	GasUsed          uint64    `json:"gas_used"`
	Timestamp        uint64    `json:"timestamp"`
	ExtraData        []byte    `json:"extra_data" ssz-max:"32"`
	BaseFeePerGas    Uint256   `json:"base_fee_per_gas" ssz-size:"32"`
	BlockHash        [32]byte  `json:"block_hash" ssz-size:"32"`
	TransactionsRoot [32]byte  `json:"transactions_root" ssz-size:"32"`
}

type SyncAggregatorSelectionData struct {
	Slot              uint64 `json:"slot"`
	SubCommitteeIndex uint64 `json:"subcommittee_index"`
}

// SyncCommitteeContribution is the Ethereum 2 sync committee contribution structure.
type SyncCommitteeContribution struct {
	Slot              uint64    `json:"slot"`
	BeaconBlockRoot   Root      `json:"beacon_block_root" ssz-size:"32"`
	SubcommitteeIndex uint64    `json:"subcommittee_index"`
	AggregationBits   []byte    `json:"aggregation_bits" ssz-size:"16"` // bitvector
	Signature         Signature `json:"signature" ssz-size:"96"`
}

type ContributionAndProof struct {
	AggregatorIndex uint64                     `json:"aggregator_index"`
	Contribution    *SyncCommitteeContribution `json:"contribution"`
	SelectionProof  Signature                  `json:"selection_proof" ssz-size:"96"`
}

type SignedContributionAndProof struct {
	Message   *ContributionAndProof `json:"message"`
	Signature Signature             `json:"signature" ssz-size:"96"`
}

type SyncCommitteeMessage struct {
	Slot           uint64    `json:"slot"`
	BlockRoot      Root      `json:"beacon_block_root" ssz-size:"32"`
	ValidatorIndex uint64    `json:"validator_index"`
	Signature      Signature `json:"signature" ssz-size:"96"`
}

type SignedAggregateAndProof struct {
	Message   *AggregateAndProof `json:"message"`
	Signature Signature          `json:"signature" ssz-size:"96"`
}

type Eth1Block struct {
	Timestamp    uint64 `json:"timestamp"`
	DepositRoot  Root   `json:"deposit_root" ssz-size:"32"`
	DepositCount uint64 `json:"deposit_count"`
}

type PowBlock struct {
	BlockHash       [32]byte `json:"block_hash" ssz-size:"32"`
	ParentHash      [32]byte `json:"parent_hash" ssz-size:"32"`
	TotalDifficulty [32]byte `json:"total_difficulty" ssz-size:"32"`
}

// Capella types

type ExecutionPayloadCapella struct {
	ParentHash    [32]byte      `ssz-size:"32" json:"parent_hash"`
	FeeRecipient  [20]byte      `ssz-size:"20" json:"fee_recipient"`
	StateRoot     [32]byte      `ssz-size:"32" json:"state_root"`
	ReceiptsRoot  [32]byte      `ssz-size:"32" json:"receipts_root"`
	LogsBloom     [256]byte     `ssz-size:"256" json:"logs_bloom"`
	PrevRandao    [32]byte      `ssz-size:"32" json:"prev_randao"`
	BlockNumber   uint64        `json:"block_number"`
	GasLimit      uint64        `json:"gas_limit"`
	GasUsed       uint64        `json:"gas_used"`
	Timestamp     uint64        `json:"timestamp"`
	ExtraData     []byte        `ssz-max:"32" json:"extra_data"`
	BaseFeePerGas Uint256       `ssz-size:"32" json:"base_fee_per_gas"`
	BlockHash     [32]byte      `ssz-size:"32" json:"block_hash"`
	Transactions  [][]byte      `ssz-max:"1048576,1073741824" ssz-size:"?,?" json:"transactions"`
	Withdrawals   []*Withdrawal `json:"withdrawals" ssz-max:"16"`
}

type ExecutionPayloadHeaderCapella struct {
	ParentHash       [32]byte  `json:"parent_hash" ssz-size:"32"`
	FeeRecipient     [20]byte  `json:"fee_recipient" ssz-size:"20"`
	StateRoot        [32]byte  `json:"state_root" ssz-size:"32"`
	ReceiptsRoot     [32]byte  `json:"receipts_root" ssz-size:"32"`
	LogsBloom        [256]byte `json:"logs_bloom" ssz-size:"256"`
	PrevRandao       [32]byte  `json:"prev_randao" ssz-size:"32"`
	BlockNumber      uint64    `json:"block_number"`
	GasLimit         uint64    `json:"gas_limit"`
	GasUsed          uint64    `json:"gas_used"`
	Timestamp        uint64    `json:"timestamp"`
	ExtraData        []byte    `json:"extra_data" ssz-max:"32"`
	BaseFeePerGas    Uint256   `json:"base_fee_per_gas" ssz-size:"32"`
	BlockHash        [32]byte  `json:"block_hash" ssz-size:"32"`
	TransactionsRoot [32]byte  `json:"transactions_root" ssz-size:"32"`
	WithdrawalRoot   [32]byte  `json:"withdrawals_root" ssz-size:"32"`
}

type BLSToExecutionChange struct {
	ValidatorIndex     uint64   `json:"validator_index"`
	FromBLSPubKey      [48]byte `json:"from_bls_pubkey" ssz-size:"48"`
	ToExecutionAddress [20]byte `json:"to_execution_address" ssz-size:"20"`
}

type HistoricalSummary struct {
	BlockSummaryRoot Root `json:"block_summary_root" ssz-size:"32"`
	StateSummaryRoot Root `json:"state_summary_root" ssz-size:"32"`
}

type SignedBLSToExecutionChange struct {
	Message   *BLSToExecutionChange `json:"message"`
	Signature Signature             `json:"signature" ssz-size:"96"`
}

type Withdrawal struct {
	Index          uint64   `json:"index"`
	ValidatorIndex uint64   `json:"validator_index"`
	Address        [20]byte `json:"address" ssz-size:"20"`
	Amount         uint64   `json:"amount"`
}

type BeaconStateCapella struct {
	GenesisTime                  uint64                         `json:"genesis_time"`
	GenesisValidatorsRoot        [32]byte                       `json:"genesis_validators_root" ssz-size:"32"`
	Slot                         uint64                         `json:"slot"`
	Fork                         *Fork                          `json:"fork"`
	LatestBlockHeader            *BeaconBlockHeader             `json:"latest_block_header"`
	BlockRoots                   [8192][32]byte                 `json:"block_roots" ssz-size:"8192,32"`
	StateRoots                   [8192][32]byte                 `json:"state_roots" ssz-size:"8192,32"`
	HistoricalRoots              [][]byte                       `json:"historical_roots" ssz-max:"16777216" ssz-size:"?,32"`
	Eth1Data                     *Eth1Data                      `json:"eth1_data"`
	Eth1DataVotes                []*Eth1Data                    `json:"eth1_data_votes" ssz-max:"2048"`
	Eth1DepositIndex             uint64                         `json:"eth1_deposit_index"`
	Validators                   []*Validator                   `json:"validators" ssz-max:"1099511627776"`
	Balances                     []uint64                       `json:"balances" ssz-max:"1099511627776"`
	RandaoMixes                  [65536][32]byte                `json:"randao_mixes" ssz-size:"65536,32"`
	Slashings                    []uint64                       `json:"slashings" ssz-size:"8192"`
	PreviousEpochParticipation   []byte                         `json:"previous_epoch_participation" ssz-max:"1099511627776"`
	CurrentEpochParticipation    []byte                         `json:"current_epoch_participation" ssz-max:"1099511627776"`
	JustificationBits            [1]byte                        `json:"justification_bits" ssz-size:"1"`
	PreviousJustifiedCheckpoint  *Checkpoint                    `json:"previous_justified_checkpoint"`
	CurrentJustifiedCheckpoint   *Checkpoint                    `json:"current_justified_checkpoint"`
	FinalizedCheckpoint          *Checkpoint                    `json:"finalized_checkpoint"`
	InactivityScores             []uint64                       `json:"inactivity_scores" ssz-max:"1099511627776"`
	CurrentSyncCommittee         *SyncCommittee                 `json:"current_sync_committee"`
	NextSyncCommittee            *SyncCommittee                 `json:"next_sync_committee"`
	LatestExecutionPayloadHeader *ExecutionPayloadHeaderCapella `json:"latest_execution_payload_header"`
	NextWithdrawalIndex          uint64                         `json:"next_withdrawal_index"`
	NextWithdrawalValidatorIndex uint64                         `json:"next_withdrawal_validator_index"`
	HistoricalSummaries          []*HistoricalSummary           `json:"historical_summaries" ssz-max:"16777216"`
}

type SignedBeaconBlockCapella struct {
	Block     *BeaconBlockCapella `json:"message"`
	Signature Signature           `json:"signature" ssz-size:"96"`
}

type BeaconBlockCapella struct {
	Slot          uint64                  `json:"slot"`
	ProposerIndex uint64                  `json:"proposer_index"`
	ParentRoot    Root                    `json:"parent_root" ssz-size:"32"`
	StateRoot     Root                    `json:"state_root" ssz-size:"32"`
	Body          *BeaconBlockBodyCapella `json:"body"`
}

type BeaconBlockBodyCapella struct {
	RandaoReveal          Signature                     `json:"randao_reveal" ssz-size:"96"`
	Eth1Data              *Eth1Data                     `json:"eth1_data"`
	Graffiti              [32]byte                      `json:"graffiti" ssz-size:"32"`
	ProposerSlashings     []*ProposerSlashing           `json:"proposer_slashings" ssz-max:"16"`
	AttesterSlashings     []*AttesterSlashing           `json:"attester_slashings" ssz-max:"2"`
	Attestations          []*Attestation                `json:"attestations" ssz-max:"128"`
	Deposits              []*Deposit                    `json:"deposits" ssz-max:"16"`
	VoluntaryExits        []*SignedVoluntaryExit        `json:"voluntary_exits" ssz-max:"16"`
	SyncAggregate         *SyncAggregate                `json:"sync_aggregate"`
	ExecutionPayload      *ExecutionPayloadCapella      `json:"execution_payload"`
	BlsToExecutionChanges []*SignedBLSToExecutionChange `json:"bls_to_execution_changes" ssz-max:"16"`
}

type LightClientHeaderCapella struct {
	Header          *BeaconBlockHeader             `json:"beacon"`
	Execution       *ExecutionPayloadHeaderCapella `json:"execution"`
	ExecutionBranch [4][32]byte                    `json:"execution_branch" ssz-size:"4,32"`
}

type LightClientBootstrapCapella struct {
	Header                     *LightClientHeaderCapella `json:"header"`
	CurrentSyncCommittee       *SyncCommittee            `json:"current_sync_committee"`
	CurrentSyncCommitteeBranch [][32]byte                `json:"current_sync_committee_branch" ssz-size:"5,32"`
}

type LightClientFinalityUpdateCapella struct {
	AttestedHeader  *LightClientHeaderCapella `json:"attested_header"`
	FinalizedHeader *LightClientHeaderCapella `json:"finalized_header"`
	FinalityBranch  [][32]byte                `json:"finality_branch" ssz-size:"6,32"`
	SyncAggregate   *SyncAggregate            `json:"sync_aggregate"`
	SignatureSlot   uint64                    `json:"signature_slot"`
}

type LightClientOptimisticUpdateCapella struct {
	AttestedHeader *LightClientHeaderCapella `json:"attested_header"`
	SyncAggregate  *SyncAggregate            `json:"sync_aggregate"`
	SignatureSlot  uint64                    `json:"signature_slot"`
}

type LightClientUpdateCapella struct {
	AttestedHeader          *LightClientHeaderCapella `json:"attested_header"`
	NextSyncCommittee       *SyncCommittee            `json:"next_sync_committee"`
	NextSyncCommitteeBranch [][32]byte                `json:"next_sync_committee_branch" ssz-size:"5,32"`
	FinalizedHeader         *LightClientHeaderCapella `json:"finalized_header"`
	FinalityBranch          [][32]byte                `json:"finality_branch" ssz-size:"6,32"`
	SyncAggregate           *SyncAggregate            `json:"sync_aggregate"`
	SignatureSlot           uint64                    `json:"signature_slot"`
}

// Deneb types

type ExecutionPayloadDeneb struct {
	ParentHash    [32]byte      `ssz-size:"32" json:"parent_hash"`
	FeeRecipient  [20]byte      `ssz-size:"20" json:"fee_recipient"`
	StateRoot     [32]byte      `ssz-size:"32" json:"state_root"`
	ReceiptsRoot  [32]byte      `ssz-size:"32" json:"receipts_root"`
	LogsBloom     [256]byte     `ssz-size:"256" json:"logs_bloom"`
	PrevRandao    [32]byte      `ssz-size:"32" json:"prev_randao"`
	BlockNumber   uint64        `json:"block_number"`
	GasLimit      uint64        `json:"gas_limit"`
	GasUsed       uint64        `json:"gas_used"`
	Timestamp     uint64        `json:"timestamp"`
	ExtraData     []byte        `ssz-max:"32" json:"extra_data"`
	BaseFeePerGas Uint256       `ssz-size:"32" json:"base_fee_per_gas"`
	BlockHash     [32]byte      `ssz-size:"32" json:"block_hash"`
	Transactions  [][]byte      `ssz-max:"1048576,1073741824" ssz-size:"?,?" json:"transactions"`
	Withdrawals   []*Withdrawal `json:"withdrawals" ssz-max:"16"`
	BlobGasUsed   uint64        `json:"blob_gas_used"`
	ExcessBlobGas uint64        `json:"excess_blob_gas"`
}

type ExecutionPayloadHeaderDeneb struct {
	ParentHash       [32]byte  `json:"parent_hash" ssz-size:"32"`
	FeeRecipient     [20]byte  `json:"fee_recipient" ssz-size:"20"`
	StateRoot        [32]byte  `json:"state_root" ssz-size:"32"`
	ReceiptsRoot     [32]byte  `json:"receipts_root" ssz-size:"32"`
	LogsBloom        [256]byte `json:"logs_bloom" ssz-size:"256"`
	PrevRandao       [32]byte  `json:"prev_randao" ssz-size:"32"`
	BlockNumber      uint64    `json:"block_number"`
	GasLimit         uint64    `json:"gas_limit"`
	GasUsed          uint64    `json:"gas_used"`
	Timestamp        uint64    `json:"timestamp"`
	ExtraData        []byte    `json:"extra_data" ssz-max:"32"`
	BaseFeePerGas    Uint256   `json:"base_fee_per_gas" ssz-size:"32"`
	BlockHash        [32]byte  `json:"block_hash" ssz-size:"32"`
	TransactionsRoot [32]byte  `json:"transactions_root" ssz-size:"32"`
	WithdrawalRoot   [32]byte  `json:"withdrawals_root" ssz-size:"32"`
	BlobGasUsed      uint64    `json:"blob_gas_used"`
	ExcessBlobGas    uint64    `json:"excess_blob_gas"`
}

type BeaconStateDeneb struct {
	GenesisTime                  uint64                       `json:"genesis_time"`
	GenesisValidatorsRoot        [32]byte                     `json:"genesis_validators_root" ssz-size:"32"`
	Slot                         uint64                       `json:"slot"`
	Fork                         *Fork                        `json:"fork"`
	LatestBlockHeader            *BeaconBlockHeader           `json:"latest_block_header"`
	BlockRoots                   [8192][32]byte               `json:"block_roots" ssz-size:"8192,32"`
	StateRoots                   [8192][32]byte               `json:"state_roots" ssz-size:"8192,32"`
	HistoricalRoots              [][]byte                     `json:"historical_roots" ssz-max:"16777216" ssz-size:"?,32"`
	Eth1Data                     *Eth1Data                    `json:"eth1_data"`
	Eth1DataVotes                []*Eth1Data                  `json:"eth1_data_votes" ssz-max:"2048"`
	Eth1DepositIndex             uint64                       `json:"eth1_deposit_index"`
	Validators                   []*Validator                 `json:"validators" ssz-max:"1099511627776"`
	Balances                     []uint64                     `json:"balances" ssz-max:"1099511627776"`
	RandaoMixes                  [65536][32]byte              `json:"randao_mixes" ssz-size:"65536,32"`
	Slashings                    []uint64                     `json:"slashings" ssz-size:"8192"`
	PreviousEpochParticipation   []byte                       `json:"previous_epoch_participation" ssz-max:"1099511627776"`
	CurrentEpochParticipation    []byte                       `json:"current_epoch_participation" ssz-max:"1099511627776"`
	JustificationBits            [1]byte                      `json:"justification_bits" ssz-size:"1"`
	PreviousJustifiedCheckpoint  *Checkpoint                  `json:"previous_justified_checkpoint"`
	CurrentJustifiedCheckpoint   *Checkpoint                  `json:"current_justified_checkpoint"`
	FinalizedCheckpoint          *Checkpoint                  `json:"finalized_checkpoint"`
	InactivityScores             []uint64                     `json:"inactivity_scores" ssz-max:"1099511627776"`
	CurrentSyncCommittee         *SyncCommittee               `json:"current_sync_committee"`
	NextSyncCommittee            *SyncCommittee               `json:"next_sync_committee"`
	LatestExecutionPayloadHeader *ExecutionPayloadHeaderDeneb `json:"latest_execution_payload_header"`
	NextWithdrawalIndex          uint64                       `json:"next_withdrawal_index"`
	NextWithdrawalValidatorIndex uint64                       `json:"next_withdrawal_validator_index"`
	HistoricalSummaries          []*HistoricalSummary         `json:"historical_summaries" ssz-max:"16777216"`
}

type SignedBeaconBlockDeneb struct {
	Block     *BeaconBlockDeneb `json:"message"`
	Signature Signature         `json:"signature" ssz-size:"96"`
}

type BeaconBlockDeneb struct {
	Slot          uint64                `json:"slot"`
	ProposerIndex uint64                `json:"proposer_index"`
	ParentRoot    Root                  `json:"parent_root" ssz-size:"32"`
	StateRoot     Root                  `json:"state_root" ssz-size:"32"`
	Body          *BeaconBlockBodyDeneb `json:"body"`
}

type BeaconBlockBodyDeneb struct {
	RandaoReveal          Signature                     `json:"randao_reveal" ssz-size:"96"`
	Eth1Data              *Eth1Data                     `json:"eth1_data"`
	Graffiti              [32]byte                      `json:"graffiti" ssz-size:"32"`
	ProposerSlashings     []*ProposerSlashing           `json:"proposer_slashings" ssz-max:"16"`
	AttesterSlashings     []*AttesterSlashing           `json:"attester_slashings" ssz-max:"2"`
	Attestations          []*Attestation                `json:"attestations" ssz-max:"128"`
	Deposits              []*Deposit                    `json:"deposits" ssz-max:"16"`
	VoluntaryExits        []*SignedVoluntaryExit        `json:"voluntary_exits" ssz-max:"16"`
	SyncAggregate         *SyncAggregate                `json:"sync_aggregate"`
	ExecutionPayload      *ExecutionPayloadDeneb        `json:"execution_payload"`
	BlsToExecutionChanges []*SignedBLSToExecutionChange `json:"bls_to_execution_changes" ssz-max:"16"`
	BlobKZGCommitments    [][48]byte                    `json:"blob_kzg_commitments" ssz-max:"4096"`
}

type LightClientHeaderDeneb struct {
	Header          *BeaconBlockHeader           `json:"beacon"`
	Execution       *ExecutionPayloadHeaderDeneb `json:"execution"`
	ExecutionBranch [4][32]byte                  `json:"execution_branch" ssz-size:"4,32"`
}

type LightClientBootstrapDeneb struct {
	Header                     *LightClientHeaderDeneb `json:"header"`
	CurrentSyncCommittee       *SyncCommittee          `json:"current_sync_committee"`
	CurrentSyncCommitteeBranch [][32]byte              `json:"current_sync_committee_branch" ssz-size:"5,32"`
}

type LightClientFinalityUpdateDeneb struct {
	AttestedHeader  *LightClientHeaderDeneb `json:"attested_header"`
	FinalizedHeader *LightClientHeaderDeneb `json:"finalized_header"`
	FinalityBranch  [][32]byte              `json:"finality_branch" ssz-size:"6,32"`
	SyncAggregate   *SyncAggregate          `json:"sync_aggregate"`
	SignatureSlot   uint64                  `json:"signature_slot"`
}

type LightClientOptimisticUpdateDeneb struct {
	AttestedHeader *LightClientHeaderDeneb `json:"attested_header"`
	SyncAggregate  *SyncAggregate          `json:"sync_aggregate"`
	SignatureSlot  uint64                  `json:"signature_slot"`
}

type LightClientUpdateDeneb struct {
	AttestedHeader          *LightClientHeaderDeneb `json:"attested_header"`
	NextSyncCommittee       *SyncCommittee          `json:"next_sync_committee"`
	NextSyncCommitteeBranch [][32]byte              `json:"next_sync_committee_branch" ssz-size:"5,32"`
	FinalizedHeader         *LightClientHeaderDeneb `json:"finalized_header"`
	FinalityBranch          [][32]byte              `json:"finality_branch" ssz-size:"6,32"`
	SyncAggregate           *SyncAggregate          `json:"sync_aggregate"`
	SignatureSlot           uint64                  `json:"signature_slot"`
}
