package db

import (
	"context"
	log "github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetDatabase(client *mongo.Client, databaseName string) (*mongo.Database, error) {
	database := client.Database(databaseName)
	return database, nil
}

func GetCollection(database *mongo.Database, collectionName string) (*mongo.Collection, error) {
	collection := database.Collection(collectionName)
	return collection, nil
}

type MongoDAL struct {
	Client *mongo.Client
	Database *mongo.Database
	LoanCollection *mongo.Collection
	DebtCollection *mongo.Collection
	InstallmentsCollection *mongo.Collection
	CollateralCollection *mongo.Collection
	DescriptorCollection *mongo.Collection
}

func NewMongoDAL(client *mongo.Client) MongoDAL {
	dal := &MongoDAL{}
	dal.Client = client
	dal.Database, _ = GetDatabase(dal.Client, "rcn")
	dal.LoanCollection, _ = GetCollection(dal.Database, "loan")
	dal.DebtCollection, _ = GetCollection(dal.Database, "debt")
	dal.InstallmentsCollection, _ = GetCollection(dal.Database, "installment")
	dal.CollateralCollection, _ = GetCollection(dal.Database, "collateral")
	dal.DescriptorCollection, _ = GetCollection(dal.Database, "descriptor")
	return *dal
}

func (self MongoDAL) GetDescriptor(id string) (*DBDescriptor, error) {
	var descriptor DBDescriptor
	filter := bson.D{{"_id", id}}
	err := self.DescriptorCollection.FindOne(context.TODO(), filter).Decode(&descriptor)
	if err != nil {
		return nil, err
	}
	return &descriptor, nil
}

func (self MongoDAL) SaveDescriptor(descriptor *DBDescriptor) error {
	_, err := self.DescriptorCollection.InsertOne(context.TODO(), descriptor)
	if err != nil {
		return err
	}
	return nil
}

func (self MongoDAL) GetLoan(id string) (*DBLoan, error) {
	var loan DBLoan
	filter := bson.D{{"_id", id}}
	err := self.LoanCollection.FindOne(context.TODO(), filter).Decode(&loan)
	if err != nil {
		return nil, err
	}

	return &loan, nil
}

func (self MongoDAL) SaveLoan(loan *DBLoan) error {
	_, err := self.LoanCollection.InsertOne(context.TODO(), loan)
	if err != nil {
		return err
	}
	return nil
}

func (self MongoDAL) UpdateLoan(loan *DBLoan) error {
	filter := bson.D{{"_id", loan.ID}}

	replacement := bson.D{
		{"open", loan.Open},
		{"approved", loan.Approved},
		{"expiration", loan.Expiration},
		{"amount", loan.Amount},
		{"cosigner", loan.Cosigner},
		{"model", loan.Model},
		{"creator", loan.Creator},
		{"oracle", loan.Oracle},
		{"borrower", loan.Borrower},
		{"loanData", loan.LoanData},
		{"created", loan.Created},
		{"currency", loan.Currency},
		{"status", loan.Status},
		{"canceled", loan.Canceled},
	}
	var replacedDocument bson.M
	err := self.LoanCollection.FindOneAndReplace(context.TODO(), filter, replacement).Decode(&replacedDocument)
	if err != nil {
		return err
	}
	return nil
}

func (self MongoDAL) DeleteLoan(loan *DBLoan) error {
	filter := bson.D{{"_id", loan.ID}}
	_, err := self.LoanCollection.DeleteOne(context.TODO(), filter)
	return err
}

func (self MongoDAL) GetDebt(id string) (*DBDebt, error) {
	var debt DBDebt
	filter := bson.D{{"_id", id}}
	err := self.DebtCollection.FindOne(context.TODO(), filter).Decode(&debt)
	if err != nil {
		return nil, err
	}
	return &debt, nil
}

func (self MongoDAL) SaveDebt(debt *DBDebt) error {
	_, err := self.DebtCollection.InsertOne(context.TODO(), debt)
	if err != nil {
		return err
	}
	return nil
}

func (self MongoDAL) UpdateDebt(debt *DBDebt) error {
	filter := bson.D{{"_id", debt.ID}}

	replacement := bson.D{
		{"error", debt.Error},
		{"balance", debt.Balance},
		{"model", debt.Model},
		{"creator", debt.Creator},
		{"oracle", debt.Oracle},
		{"created", debt.Model},
		{"owner", debt.Owner},
	}
	var replacedDocument bson.M
	err := self.DebtCollection.FindOneAndReplace(context.TODO(), filter, replacement).Decode(&replacedDocument)
	if err != nil {
		return err
	}
	return nil
}

func (self MongoDAL) DeleteDebt(debt *DBDebt) error {
	filter := bson.D{{"_id", debt.ID}}
	_, err := self.DebtCollection.DeleteOne(context.TODO(), filter)
	return err
}

func (self MongoDAL) GetInstallment(id string) (*DBInstallment, error) {
	var installment DBInstallment
	filter := bson.D{{"_id", id}}
	err := self.InstallmentsCollection.FindOne(context.TODO(), filter).Decode(&installment)
	if err != nil {
		return nil, err
	}

	return &installment, nil
}

func (self MongoDAL) SaveInstallment(installment *DBInstallment) error {
	_, err := self.InstallmentsCollection.InsertOne(context.TODO(), installment)
	if err != nil {
		return err
	}
	return nil
}

func (self MongoDAL) UpdateInstallment(installment *DBInstallment) error {
	filter := bson.D{{"_id", installment.ID}}

	replacement := bson.D{
		{"config", installment.Config},
		{"state", installment.State},
	}
	var replacedDocument bson.M
	err := self.InstallmentsCollection.FindOneAndReplace(context.TODO(), filter, replacement).Decode(&replacedDocument)
	if err != nil {
		return err
	}
	return nil
}

func (self MongoDAL) DeleteInstallment (installment *DBInstallment) error {
	filter := bson.D{{"_id", installment.ID}}
	_, err := self.InstallmentsCollection.DeleteOne(context.TODO(), filter)
	return err
}

func (self MongoDAL) GetCollateral (id string) (*DBCollateral, error) {
	var collateral DBCollateral
	log.Info("GET COLLATERAL:", id)
	filter := bson.D{{"_id", id}}
	err := self.CollateralCollection.FindOne(context.TODO(), filter).Decode(&collateral)
	if err != nil {
		return nil, err
	}

	return &collateral, nil
}

func (self MongoDAL) GetCollaterals (debtID string) ([]DBCollateral, error) {
	var collaterals []DBCollateral
	filter := bson.D{{"debt_id", debtID}}

	// Here's an array in which you can store the decoded documents

	// Passing bson.D{{}} as the filter matches all documents in the collection
	cur, err := self.CollateralCollection.Find(context.TODO(), filter)
	if err != nil {
		return collaterals, err
	}

	// Finding multiple documents returns a cursor
	// Iterating through the cursor allows us to decode documents one at a time
	for cur.Next(context.TODO()) {

		// create a value into which the single document can be decoded
		var elem DBCollateral
		err := cur.Decode(&elem)
		if err != nil {
			log.Fatal(err)
		}

		collaterals = append(collaterals, elem)
	}

	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}

	// Close the cursor once finished
	cur.Close(context.TODO())

	return collaterals, nil
}

func (self MongoDAL) SaveCollateral(collateral *DBCollateral) error {
	_, err := self.CollateralCollection.InsertOne(context.TODO(), collateral)
	if err != nil {
		return err
	}
	return nil
}

func (self MongoDAL) UpdateCollateral(collateral *DBCollateral) error {
	filter := bson.D{{"_id", collateral.ID}}

	replacement := bson.D{
		{"debt_id", collateral.DebtID},
		{"oracle", collateral.Oracle},
		{"token", collateral.Token},
		{"liquidation_ratio", collateral.LiquidationRatio},
		{"balance_ratio", collateral.BalanceRatio},
		{"burn_fee", collateral.BurnFee},
		{"reward_fee", collateral.RewardFee},
		{"owner", collateral.Owner},
		{"amount", collateral.Amount},
		{"status", collateral.Status},
	}
	var replacedDocument bson.M
	err := self.CollateralCollection.FindOneAndReplace(context.TODO(), filter, replacement).Decode(&replacedDocument)
	if err != nil {
		return err
	}
	return nil
}

func (self MongoDAL) DeleteCollateral(collateral *DBCollateral) error {
	filter := bson.D{{"_id", collateral.ID}}
	_, err := self.CollateralCollection.DeleteOne(context.TODO(), filter)
	return err
}