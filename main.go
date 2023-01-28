package main

import (
	"github.com/julienschmidt/httprouter"
	"gorm_test_6/advanced_topics/serializer"
	"gorm_test_6/app"
	"gorm_test_6/tutorials/security"
	"net/http"
)

func main() {
	//dsn := "host=localhost user=ijonk password=lionstar dbname=gorm_test_6 port=5432 sslmode=disable TimeZone=Asia/Jakarta"
	//db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	//
	//if err != nil {
	//	fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
	//	os.Exit(1)
	//}

	db := app.NewDB()

	// Migrate the schema
	//db.AutoMigrate(&model.User{})
	//db.AutoMigrate(&model.CreditCard{})
	//db.AutoMigrate(&model.User2{})
	//db.AutoMigrate(&model.Product{})
	//db.AutoMigrate(&model.User3{})
	//db.AutoMigrate(&model.Company{})
	//db.AutoMigrate(&model.User4{})
	//db.AutoMigrate(&model.CreditCard4{})
	//db.AutoMigrate(&model.User5{})
	//db.AutoMigrate(&model.CreditCard5{})
	//db.AutoMigrate(&model.User6{})
	//db.AutoMigrate(&model.Language{})

	/*
		CRUD Interface - Create
	*/
	//create.CreateRecord(db)
	//create.CreateRecord2(db)
	//create.BatchInsert(db)
	//create.BatchSize(db)
	//create.CreateFromMap1(db)
	//create.CreateFromMap2(db)
	//create.CreateWithAssociations(db)

	/*
		CRUD Interface - Query
	*/
	//query.RetrievingSingleObject1(db)
	//query.RetrievingSingleObject2(db)
	//query.RetrievingSingleObject3(db)
	//query.RetrievingObjectsWithPrimaryKey(db)
	//query.RetrievingAllObjects(db)
	//query.ConditionsWhere(db)
	//query.Order(db)
	//query.LimitOffset(db)
	//query.GroupByHaving(db)
	//query.Distinct(db)
	//query.Scan(db)

	/*
		CRUD Interface - Advanced Query
	*/
	//advanced_query.SmartSelectFields(db)
	//advanced_query.FirstOrInit(db)
	//advanced_query.FirstOrCreate(db)
	//advanced_query.OptimizerIndexHints(db)
	//advanced_query.Iteration(db)
	//advanced_query.Pluck(db)
	//advanced_query.Scopes(db)
	//advanced_query.Count(db)

	/*
		CRUD Interface - Update
	*/
	//update.SaveAllFields(db)
	//update.UpdateSingleColumn(db)
	//update.UpdatesMultipleColumns(db)
	//update.UpdateSelectedFields(db)
	//update.TesUpdateHooks(db)

	/*
		CRUD Interface - Delete
	*/
	//delete2.DeleteRecord(db)
	//delete2.DeleteWithPrimaryKey(db)
	//delete2.SoftDelete(db)

	/*
		CRUD Interface - Raw SQL & SQL Builder
	*/
	//sql_builder.RawSql(db)
	//sql_builder.NamedArgument(db)
	//sql_builder.DryRunMode(db)
	//sql := sql_builder.ToSQL(db)
	//fmt.Println(sql)
	//sql_builder.RowRows(db)

	/*
		Associations - Belongs To
	*/
	//router := httprouter.New()
	//router.GET("/api/belongs-to", belongs_to.BelongsTo)
	//
	//server := http.Server{
	//	Addr:    "localhost:3000",
	//	Handler: router,
	//}
	//server.ListenAndServe()

	/*
		Associations - Has One
	*/
	//router := httprouter.New()
	//router.GET("/api/has-one", has_one.HasOne)
	//
	//server := http.Server{
	//	Addr:    "localhost:3000",
	//	Handler: router,
	//}
	//server.ListenAndServe()

	/*
		Associations - Has Many
	*/
	//router := httprouter.New()
	//router.GET("/api/has-many", has_many.HasMany)
	//
	//server := http.Server{
	//	Addr:    "localhost:3000",
	//	Handler: router,
	//}
	//server.ListenAndServe()

	/*
		Associations - Many To Many
	*/
	//router := httprouter.New()
	//router.GET("/api/many-to-many", many_to_many.ManyToMany)
	//
	//server := http.Server{
	//	Addr:    "localhost:3000",
	//	Handler: router,
	//}
	//server.ListenAndServe()

	/*
		Associations - Associations Mode
	*/
	//router := httprouter.New()
	//router.GET("/api/auto-create-update", associations_mode.AutoCreateUpdate)
	//router.GET("/api/associations-mode", associations_mode.AssociationMode)
	//router.GET("/api/find-associations", associations_mode.FindAssociations)
	//router.GET("/api/delete-with-select", associations_mode.DeleteWithSelect)
	//
	//server := http.Server{
	//	Addr:    "localhost:3000",
	//	Handler: router,
	//}
	//server.ListenAndServe()

	/*
		Associations - Preload with conditions
	*/
	//router := httprouter.New()
	//router.GET("/api/preload", preloading.Preload)
	//router.GET("/api/preload-all", preloading.PreloadAll)
	//router.GET("/api/preload-with-conditions", preloading.PreloadWithConditions)
	//router.GET("/api/custom-preloading-sql", preloading.CustomPreloadingSql)
	//
	//server := http.Server{
	//	Addr:    "localhost:3000",
	//	Handler: router,
	//}
	//server.ListenAndServe()

	/*
		Tutorials - Context
	*/
	//context.ContinuousSessionMode(db)

	/*
		Tutorials - Session
	*/
	//session.DryRun(db)

	/*
		Tutorials - Transactions
	*/
	//transactions.Transaction(db)
	//transactions.NestedTransactions(db)
	//transactions.ControlTransactionManually(db)
	//transactions.SpecificExample(db)

	/*
		Tutorials - Migration
	*/
	//migration.CurrentDatabase(db)
	//migration.Tables(db)
	//migration.Columns(db)

	/*
		Tutorials - Scopes
	*/
	//scopes.Query(db)

	/*
		Tutorials - Conventions
	*/
	//conventions.PluralizedTableName(db)
	//conventions.TemporarilySpecifyName(db)
	//conventions.ColumnName(db)
	//conventions.TimestampTracking(db)

	/*
		Advanced Topics - Serializer
	*/
	router := httprouter.New()
	router.GET("/api/serializer", serializer.Serializer)

	server := http.Server{
		Addr:    "localhost:3000",
		Handler: router,
	}
	server.ListenAndServe()

	/*
		Tutorials - Indexes
	*/
	//indexes.IndexTag(db)

	/*
		Tutorials - Constraints
	*/
	//constraints.ForeignKeyConstraint(db)

	/*
		Tutorials - Composite Primary Key
	*/
	//composite_primary_key.CompositePrimaryKey(db)

	/*
		Tutorials - Security
	*/
	security.QueryCondition(db)

}
