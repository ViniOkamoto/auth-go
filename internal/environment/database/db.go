package database

// type DBConfing interface {
// 	AddError(err error) error
// 	Assign(attrs ...interface{}) (tx *gorm.DB)
// 	Association(column string) *gorm.Association
// 	Attrs(attrs ...interface{}) (tx *gorm.DB)
// 	AutoMigrate(dst ...interface{}) error
// 	Begin(opts ...*sql.TxOptions) *gorm.DB
// 	Clauses(conds ...clause.Expression) (tx *gorm.DB)
// 	Commit() *gorm.DB
// 	Connection(fc func(tx *gorm.DB) error) (err error)
// 	Count(count *int64) (tx *gorm.DB)
// 	Create(value interface{}) (tx *gorm.DB)
// 	CreateInBatches(value interface{}, batchSize int) (tx *gorm.DB)
// 	DB() (*sql.DB, error)
// 	Debug() (tx *gorm.DB)
// 	Delete(value interface{}, conds ...interface{}) (tx *gorm.DB)
// 	Distinct(args ...interface{}) (tx *gorm.DB)
// 	Exec(sql string, values ...interface{}) (tx *gorm.DB)
// 	Find(dest interface{}, conds ...interface{}) (tx *gorm.DB)
// 	FindInBatches(dest interface{}, batchSize int, fc func(tx *gorm.DB, batch int) error) *gorm.DB
// 	First(dest interface{}, conds ...interface{}) (tx *gorm.DB)
// 	FirstOrCreate(dest interface{}, conds ...interface{}) (tx *gorm.DB)
// 	FirstOrInit(dest interface{}, conds ...interface{}) (tx *gorm.DB)
// 	Get(key string) (interface{}, bool)
// 	Group(name string) (tx *gorm.DB)
// 	Having(query interface{}, args ...interface{}) (tx *gorm.DB)
// 	InnerJoins(query string, args ...interface{}) (tx *gorm.DB)
// 	InstanceGet(key string) (interface{}, bool)
// 	InstanceSet(key string, value interface{}) *gorm.DB
// 	Joins(query string, args ...interface{}) (tx *gorm.DB)
// 	Last(dest interface{}, conds ...interface{}) (tx *gorm.DB)
// 	Limit(limit int) (tx *gorm.DB)
// 	Migrator() gorm.Migrator
// 	Model(value interface{}) (tx *gorm.DB)
// 	Not(query interface{}, args ...interface{}) (tx *gorm.DB)
// 	Offset(offset int) (tx *gorm.DB)
// 	Omit(columns ...string) (tx *gorm.DB)
// 	Or(query interface{}, args ...interface{}) (tx *gorm.DB)
// 	Order(value interface{}) (tx *gorm.DB)
// 	Pluck(column string, dest interface{}) (tx *gorm.DB)
// 	Preload(query string, args ...interface{}) (tx *gorm.DB)
// 	Raw(sql string, values ...interface{}) (tx *gorm.DB)
// 	Rollback() *gorm.DB
// 	RollbackTo(name string) *gorm.DB
// 	Row() *sql.Row
// 	Rows() (*sql.Rows, error)
// 	Save(value interface{}) (tx *gorm.DB)
// 	SavePoint(name string) *gorm.DB
// 	Scan(dest interface{}) (tx *gorm.DB)
// 	ScanRows(rows *sql.Rows, dest interface{}) error
// 	Scopes(funcs ...func(*gorm.DB) *gorm.DB) (tx *gorm.DB)
// 	Select(query interface{}, args ...interface{}) (tx *gorm.DB)
// 	Session(config *gorm.Session) *gorm.DB
// 	Set(key string, value interface{}) *gorm.DB
// 	SetupJoinTable(model interface{}, field string, joinTable interface{}) error
// 	Table(name string, args ...interface{}) (tx *gorm.DB)
// 	Take(dest interface{}, conds ...interface{}) (tx *gorm.DB)
// 	ToSQL(queryFn func(tx *gorm.DB) *gorm.DB) string
// 	Transaction(fc func(tx *gorm.DB) error, opts ...*sql.TxOptions) (err error)
// 	Unscoped() (tx *gorm.DB)
// 	Update(column string, value interface{}) (tx *gorm.DB)
// 	UpdateColumn(column string, value interface{}) (tx *gorm.DB)
// 	UpdateColumns(values interface{}) (tx *gorm.DB)
// 	Updates(values interface{}) (tx *gorm.DB)
// 	Use(plugin gorm.Plugin) error
// 	Where(query interface{}, args ...interface{}) (tx *gorm.DB)
// 	WithContext(ctx context.Context) *gorm.DB
// }

// type DBConfigImpl struct {
// 	DB *gorm.DB
// }

// func CreateDBConfig(db *gorm.DB) DBConfing {
// 	return &DBConfigImpl{DB: db}
// }

// func (d *DBConfigImpl) AutoMigrate(values ...interface{}) error {
// 	return d.DB.AutoMigrate(values...)
// }

// func (d *DBConfigImpl) Create(value interface{}) error {
// 	return d.DB.Create(value).Error
// }

// func (d *DBConfigImpl) Find(value interface{}) error {
// 	return d.DB.Find(value).Error
// }

// func (d *DBConfigImpl) First(value interface{}, conds ...interface{}) error {
// 	return d.DB.First(value, conds...).Error
// }

// func (d *DBConfigImpl) Save(value interface{}) error {
// 	return d.DB.Save(value).Error
// }

// func (d *DBConfigImpl) Delete(value interface{}) error {
// 	return d.DB.Delete(value).Error
// }

// func (d *DBConfigImpl) FindAll(value interface{}) error {
// 	return d.DB.Find(value).Error
// }

// func (d *DBConfigImpl) FindById(value interface{}, id uint) error {
// 	return d.DB.First(value, id).Error
// }

// func (d *DBConfigImpl) Update(value interface{}, id uint) error {
// 	return d.DB.Save(value).Error
// }

// func (d *DBConfigImpl) Connection(fc func(tx *gorm.DB) error) error {
// 	return d.DB.Connection(fc)
// }

// func (d *DBConfigImpl) Begin(opts ...*sql.TxOptions) *gorm.DB {
// 	return d.DB.Begin(opts...)
// }

// func (d *DBConfigImpl) Commit() *gorm.DB {
// 	return d.DB.Commit()
// }

// func (d *DBConfigImpl) Rollback() *gorm.DB {
// 	return d.DB.Rollback()
// }

// func (d *DBConfigImpl) RollbackTo(name string) *gorm.DB {
// 	return d.DB.RollbackTo(name)
// }

// func (d *DBConfigImpl) Exec(sql string, values ...interface{}) *gorm.DB {
// 	return d.DB.Exec(sql, values...)
// }

// func (d *DBConfigImpl) Raw(sql string, values ...interface{}) *gorm.DB {
// 	return d.DB.Raw(sql, values...)
// }

// func (d *DBConfigImpl) Table(name string, args ...interface{}) *gorm.DB {
// 	return d.DB.Table(name, args...)
// }

// func (d *DBConfigImpl) Where(query interface{}, args ...interface{}) *gorm.DB {
// 	return d.DB.Where(query, args...)
// }

// func (d *DBConfigImpl) Scopes(funcs ...func(*gorm.DB) *gorm.DB) *gorm.DB {
// 	return d.DB.Scopes(funcs...)
// }

// func (d *DBConfigImpl) Session(config *gorm.Session) *gorm.DB {
// 	return d.DB.Session(config)
// }

// func (d *DBConfigImpl) WithContext(ctx context.Context) *gorm.DB {
// 	return d.DB.WithContext(ctx)
// }

// func (d *DBConfigImpl) Assign(attrs ...interface{}) *gorm.DB {
// 	return d.DB.Assign(attrs...)
// }

// func (d *DBConfigImpl) Omit(columns ...string) *gorm.DB {
// 	return d.DB.Omit(columns...)
// }

// func (d *DBConfigImpl) Select(query interface{}, args ...interface{}) *gorm.DB {
// 	return d.DB.Select(query, args...)
// }

// func (d *DBConfigImpl) Clauses(conds ...clause.Expression) *gorm.DB {
// 	return d.DB.Clauses(conds...)
// }

// func (d *DBConfigImpl) Not(query interface{}, args ...interface{}) *gorm.DB {
// 	return d.DB.Not(query, args...)
// }

// func (d *DBConfigImpl) Or(query interface{}, args ...interface{}) *gorm.DB {
// 	return d.DB.Or(query, args...)
// }

// func (d *DBConfigImpl) Group(name string) *gorm.DB {
// 	return d.DB.Group(name)
// }

// func (d *DBConfigImpl) Having(query interface{}, args ...interface{}) *gorm.DB {
// 	return d.DB.Having(query, args...)
// }

// func (d *DBConfigImpl) Joins(query string, args ...interface{}) *gorm.DB {
// 	return d.DB.Joins(query, args...)
// }

// func (d *DBConfigImpl) InnerJoins(query string, args ...interface{}) *gorm.DB {
// 	return d.DB.InnerJoins(query, args...)
// }

// func (d *DBConfigImpl) Preload(query string, args ...interface{}) *gorm.DB {
// 	return d.DB.Preload(query, args...)
// }

// func (d *DBConfigImpl) FirstOrCreate(dest interface{}, conds ...interface{}) *gorm.DB {
// 	return d.DB.FirstOrCreate(dest, conds...)
// }

// func (d *DBConfigImpl) FirstOrInit(dest interface{}, conds ...interface{}) *gorm.DB {
// 	return d.DB.FirstOrInit(dest, conds...)
// }

// func (d *DBConfigImpl) UpdateColumn(column string, value interface{}) *gorm.DB {
// 	return d.DB.UpdateColumn(column, value)
// }

// func (d *DBConfigImpl) UpdateColumns(values interface{}) *gorm.DB {
// 	return d.DB.UpdateColumns(values)
// }

// func (d *DBConfigImpl) Updates(values interface{}) *gorm.DB {
// 	return d.DB.Updates(values)
// }
