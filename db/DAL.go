package db

type DBDescriptor struct {
	ID string `bson:"_id"`
	FirstObligation string `bson:first_obligation`
	TotalObligation string `bson:total_obligation`
	Duration string `bson:duration`
	InterestRate string `bson:interest_rate`
	PunitiveInterestRate string `bson:punitive_interest_rate`
	Frequency string `bson:frequency`
	Installments string `bson:installments`
}

type DBLoan struct {
	ID string `bson:"_id"`
	Open bool `bson:open`
	Approved bool `bson:approved`
	Expiration string `bson:expiration`
	Amount string `bson:amount`
	Cosigner string `bson:cosigner`
	Model string `bson:model`
	Creator string `bson:creator`
	Oracle string `bson:oracle`
	Borrower string `bson:borrower`
	LoanData string `bson:loanData`
	Created string `bson:created`
	Currency string `bson:currency`
	Status string `bson:status`
	Canceled bool `bson:canceled`
}

type DBDebt struct {
	ID string `bson:"_id"`
	Error bool `bson:error`
	Balance string `bson:balance`
	Model string `bson:model`
	Creator string `bson:creator`
	Oracle string `bson:oracle`
	Created string `bson:created`
	Owner string `bson:owner`
}

type State struct {
	Status string `bson:status`
	Clock string `bson:clock`
	LastPayment string `bson:last_payment`
	Paid string `bson:paid`
	PaidBase string `bson:paid_base`
	Interest string `bson:interest`
}

type Config struct {
	Installments string `bson:installments`
	TimeUnit string `bson:time_unit`
	Duration string `bson:duration`
	LentTime string `bson:lent_time`
	Cuota string `bson:cuota`
	InterestRate string `bson:interest_rate`
}
type DBInstallment struct {
	ID string `bson:"_id"`
	Config *Config `bson:config`
	State *State `bson:state`
}

type DBCollateral struct {
	ID string `bson:"_id"`
	DebtID string `bson:"debt_id"`
	Oracle string `bson:oracle`
	Token string `bson:token`
	LiquidationRatio string `bson:liquidation_ratio`
	BalanceRatio string `bson:balance_ratio`
	BurnFee string `bson:burn_fee`
	RewardFee string `bson:reward_fee`
	Owner string `bson:owner`
	Amount string `bson:amount`
	Status string `bson:status`
}

type DAL interface {
	// LOAN
	GetLoan(id string) (*DBLoan, error)
	SaveLoan(loan *DBLoan) error
	UpdateLoan(loan *DBLoan) error
	DeleteLoan(loan *DBLoan) error

	// DEBT
	GetDebt(id string) (*DBDebt, error)
	SaveDebt(debt *DBDebt) error
	UpdateDebt(debt *DBDebt) error
	DeleteDebt(debt *DBDebt) error

	// INSTALLMENTS
	GetInstallment(id string) (*DBInstallment, error)
	SaveInstallment(installment *DBInstallment) error
	UpdateInstallment(installment *DBInstallment) error
	DeleteInstallment(installment *DBInstallment) error

	// Collateral
	GetCollateral(id string) (*DBCollateral, error)
	GetCollaterals(debtID string) ([]DBCollateral, error)
	SaveCollateral(collateral *DBCollateral) error
	UpdateCollateral(collateral *DBCollateral) error
	DeleteCollateral(collateral *DBCollateral) error

	// Descriptor
	GetDescriptor(id string) (*DBDescriptor, error)
	SaveDescriptor(descriptor *DBDescriptor) error
}
