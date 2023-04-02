package options

import "go.mongodb.org/mongo-driver/mongo/options"

var RawAggregate = options.Aggregate
var RawChangeStream = options.ChangeStream
var RawClient = options.Client
var RawCreateCollection = options.CreateCollection
var RawIndex = options.Index
var RawInsertOne = options.InsertOne
var RawInsertMany = options.InsertMany
var RawFind = options.Find
var RawDatabase = options.Database
var RawRemove = options.Delete
var RawReplace = options.Replace
var RawRunCommand = options.RunCmd
var RawSession = options.Session
var RawTransaction = options.Transaction
var RawUpdate = options.Update

func Aggregate(opt *options.AggregateOptions) AggregateOptions {
	if opt == nil {
		opt = options.Aggregate()
	}
	return AggregateOptions{
		AggregateOptions: opt,
	}
}

func ChangeStream(opt *options.ChangeStreamOptions) ChangeStreamOptions {
	if opt == nil {
		opt = options.ChangeStream()
	}
	return ChangeStreamOptions{
		ChangeStreamOptions: opt,
	}
}

func Client(opt *options.ClientOptions) ClientOptions {
	if opt == nil {
		opt = options.Client()
	}
	return ClientOptions{
		ClientOptions: opt,
	}
}

func CreateCollection(opt *options.CreateCollectionOptions) CreateCollectionOptions {
	if opt == nil {
		opt = options.CreateCollection()
	}
	return CreateCollectionOptions{
		CreateCollectionOptions: opt,
	}
}

func Database(opt *options.DatabaseOptions) DatabaseOptions {
	if opt == nil {
		opt = options.Database()
	}
	return DatabaseOptions{
		DatabaseOptions: opt,
	}
}

func Index(opt *options.IndexOptions, keys ...string) IndexModel {
	if opt == nil {
		opt = options.Index()
	}
	return IndexModel{
		Key:          keys,
		IndexOptions: opt,
	}
}

func InsertOne(opt *options.InsertOneOptions, hook ...interface{}) InsertOneOptions {
	if opt == nil {
		opt = options.InsertOne()
	}
	var h interface{}
	if len(hook) > 0 {
		h = hook[0]
	}
	return InsertOneOptions{
		InsertHook:       h,
		InsertOneOptions: opt,
	}
}

func InsertMany(opt *options.InsertManyOptions, hook ...interface{}) InsertManyOptions {
	if opt == nil {
		opt = options.InsertMany()
	}
	var h interface{}
	if len(hook) > 0 {
		h = hook[0]
	}
	return InsertManyOptions{
		InsertHook:        h,
		InsertManyOptions: opt,
	}
}

func Find(hook ...interface{}) FindOptions {
	var h interface{}
	if len(hook) > 0 {
		h = hook[0]
	}
	return FindOptions{
		QueryHook: h,
	}
}

func Remove(opt *options.DeleteOptions, hook ...interface{}) RemoveOptions {
	if opt == nil {
		opt = options.Delete()
	}
	var h interface{}
	if len(hook) > 0 {
		h = hook[0]
	}
	return RemoveOptions{
		RemoveHook:    h,
		DeleteOptions: opt,
	}
}

func Replace(opt *options.ReplaceOptions, hook ...interface{}) ReplaceOptions {
	if opt == nil {
		opt = options.Replace()
	}
	var h interface{}
	if len(hook) > 0 {
		h = hook[0]
	}
	return ReplaceOptions{
		UpdateHook:     h,
		ReplaceOptions: opt,
	}
}

func RunCommand(opt *options.RunCmdOptions) RunCommandOptions {
	if opt == nil {
		opt = options.RunCmd()
	}
	return RunCommandOptions{
		RunCmdOptions: opt,
	}
}

func Session(opt *options.SessionOptions) SessionOptions {
	if opt == nil {
		opt = options.Session()
	}
	return SessionOptions{
		SessionOptions: opt,
	}
}

func Transaction(opt *options.TransactionOptions) TransactionOptions {
	if opt == nil {
		opt = options.Transaction()
	}
	return TransactionOptions{
		TransactionOptions: opt,
	}
}

func Update(opt *options.UpdateOptions, hook ...interface{}) UpdateOptions {
	if opt == nil {
		opt = options.Update()
	}
	var h interface{}
	if len(hook) > 0 {
		h = hook[0]
	}
	return UpdateOptions{
		UpdateHook:    h,
		UpdateOptions: opt,
	}
}

func Upsert(opt *options.ReplaceOptions, hook ...interface{}) UpsertOptions {
	if opt == nil {
		opt = options.Replace()
	}
	var h interface{}
	if len(hook) > 0 {
		h = hook[0]
	}
	return UpsertOptions{
		UpsertHook:     h,
		ReplaceOptions: opt,
	}
}
