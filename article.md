アーキテクチャのアの字も知らない駆け出しエンジニアがドメイン駆動開発における依存性の逆転をgolangで実装しながら理解してみた

こんにちは！例のウイルスで大学の休みが延期されて、引きこもりながらますます楽しいプログラミングライフを送っているとさです。

今回は、ドメイン駆動開発における「アプリケーションサービス」と「リポジトリ」の部分の実装をしながら、
学習していきたいと思います。機能を分割して、コードをわかりやすくすることを目標にやっていきます。

# アプリケーションサービス

A君
アプリケーションサービスは、ユーザーのユースケースの機能を実装する部分になります。

B君
ユースケースってなんぞや？？

A君
ユースケースっていうのは、ユーザーがシステムに対してできること
例えば、よくあるSNSで言えば、「ユーザーデータを登録」したり「ユーザー情報を更新」したり「ツイート」したり
なんてものがユースケースに当たるよ。

B君
へ〜、つまりはユーザーがシステムに対してできることって感じだな。

A君
そうそう、その機能を実装する部分を今回はアプリケーションサービスって読んでいるんだ！
じゃあ早速ディレクトリを作成していこう！

ddd_sample
    - application

このapplicationディレクトリ内にあとでコードを書いていくよ！

B君
ここで、ユースケースを実装していくんだね。

# リポジトリ

A君
リポジトリでは、データを永続化する処理を担当する部分になります
データを永続化するためには、データベースに保存したり、ファイルに書き出して保存したり
しなくちゃいけないよね。そのような機能をリポジトリが担当するよ。

B君
なるほど、リポジトリはデータベースとのやり取り担当なんだね

A君
そうだね！具体的にはデータベースからデータを取ってきたり、受け取ったデータを元に保存したり
更新したりする役割を担うよ。
さっきのサンプルにリポジトリも作成していこうか！

B君
了解！mkdir ...っと

ddd_sample
    - application
    - repository

できた！ここに、データベースに関連する処理を書くんだね！

A君
そう。applicationでユーザーがユーザーを登録したり、変更したりって処理をrepositoryさんに
お願いしてデータベースに反映してもらうんだ。じゃあ、どんどん実装していきましょう！

B君
ラジャー!

# 依存性逆転の原則

B君
めんどいな〜

A君
どうした？？

いや、言われた通りにアプリケーションサービスとリポジトリに機能を分けて作成してみたんですよ

A君
うんうん

B君
それで、別々にテストできたら便利やろうなって思ってテストコード書こうとしたんだけど、、、
データベースとの接続を担当するリポジトリの部分はわかるんやけど、アプリケーションサービスの部分まで
データベースを用意しないとテストできないんですよ。。。

A君
なるほど

B君
機能を分けたのに、、別々にテストできないなんて、この構造にした意味あるの？？って思いまして...

A君
いい質問だね！今の君の書いたコードを一緒に様子を見てみようか

はい！

A君
ユーザーの登録処理を作ってみたんだね！

B君
そうです。ユーザーが名前を入力してそのデータを登録する処理を作りました！

A君
うんうん、言われた通りにapplicationでユーザーの名前を受け取って、データベースに登録する部分は
きちんと、リポジトリに分割できてるね！

ddd_sample
    - application
        - user_service.go
    - repository
        - user_repository.go
    - domain
        - model
            - user.go

```go
type User struct {
	ID         uint64 `gorm:"primary_key"`
	Name       string
	Updatetime time.Time
	Createtime time.Time
}
```

```go
// リポジトリの実装（データベースの処理を書くよ）
//　構造体を定義（多言語で言うクラスのようなもの)
type UserRepository struct {
    // DBインスタンスのポインタを保存
    // このデータベースですよって指してる
	DB *gorm.DB
}

// この構造体の実態を返す関数
// このデータベースを使用してね！ってきたらそれをセットしてUserRepositoryを返す
func NewUserRepository(DB *gorm.DB) repository.UserRepository {
	return &UserRepositoryt{
		DB: DB,
	}
}

// 構造体のメソッド
// userのデータを受け取って、データベースに新規保存している
// 返り値は(*model.User, error)でUserとエラーを返している
func (r *UserRepository) Save(user *model.User) (*model.User, error) {
	r.DB.Create(&user)
	return user, nil
}

```

```go
// ユーザーサービスの実装
// 構造体を定義
type UserService struct {
    //　リポジトリーを受け取っている
	UserRepository repository.UserRepository
}

// リポジトリーをセット
func NewUserService(repository repository.UserRepository) UserService {
	return UserService{UserRepository: repository}
}

// ユーザー登録の処理を書くよ！
func (UserService *UserService) CreateUser(name string) (*model.User, error) {
    // 名前を受け取って、user構造体を作成したら
	user := model.User{
		Name:       name,
		Createtime: time.Now(),
		Updatetime: time.Now(),
    }
    // データベースの処理はリポジトリに任せるので
    // リポジトリを使用して保存します
	return UserService.UserRepository.Save(&user)
}
```
 ※ golangのormマッパーであるgormを使用しています。


A君
 今のこのアプリの依存関係を図にしてみると
 依存関係っていうのは、あるオブジェクトから別のオブジェクトを参照するときに
 発生するんだ。
 今回の場合だとUserServiceがUserRepositoryに依存しているということができるよね！
[図1]application -> repository

B君
そうですね！でも、リポジトリにデータベースを保存する処理を書いている
のなら依存しなくちゃ、データを保存できないですよ？

A君
そうだね、ただこの実装だとリポジトリだけではなくアプリケーションサービスも
データベースのような特定の技術基盤に結びついてしまうよね、例えばデータベースがMySQLなのか
NoSQLなのか、はたまたローカルに保存するのかでリポジトリだけではなくアプリケーション
サービスまで書き換えなくてはいけなくなってしまうよね。

そこでこの「依存関係逆転の法則」が出てくるんだ。

その法則の一つに「抽象は実装の詳細に依存してはならない。実装の詳細に抽象が依存すべきである」

- ドメイン駆動設計入門 p165 より参照

ってのがあるのだけど

B君
うう、「抽象は実装の詳細に依存してはならない。実装の詳細に抽象が依存すべきである」って言葉が抽象的すぎてよくわからない...

A君
まあまあ、じゃあ一緒に実装しながら学習していこうか
さっきの図をもう一度出すよ！

図１

これにinterfaceを噛ませてあげると

図2


の図2ようにすることができるんだ！IUserRepository（抽象）がInterfaceでUserRepository（実装の詳細）に具体的な実装をするよ
UserService（実装の詳細）の依存先（オブジェクトの参照先）をIUserRepotory（抽象）に変更してみると
依存関係が全部IUserRepository（抽象）に向けることができて「依存関係逆転の法則」を満たすことができたね！

ディレクトリ構造

ddd_sample
    - application
        - user_service.go
    - repository
        - user_repository_interface.go <- わかりやすくするために名称変更
    - infrastructure
        - datastore
            - user_repository.go <- 具体的な実装はこちらに移動
    - domain
        - model
            - user.go

user_repository_interface.go
```go
// Interface
// ここでは具体的な実装はしないが
// ここに含まれている四つの関数を実装してくださいね〜って指定しています
type IUserRepository interface {
	GetByID(id uint64) (model.User, error)
	Save(user *model.User) (*model.User, error)
	Update(user *model.User) error
	Delete(id uint64) error
}
```

user_repository.go
```go
// Impliment ... 実装する
// UserRepositoryを実装しています
// IUserRepositoryのinterfaceで指定された四つの関数を具体的に実装しています。
type UserRepositoryImpliment struct {
	DB *gorm.DB
}

func NewUserRepositoryImpliment(DB *gorm.DB) repository.IUserRepository {
	return &UserRepositoryImpliment{
		DB: DB,
	}
}

func (r *UserRepositoryImpliment) GetByID(id uint64) (model.User, error) {
	var user model.User
	r.DB.Where("id = ?", id).Find(&user)
	return user, nil
}

func (r *UserRepositoryImpliment) Save(user *model.User) (*model.User, error) {
	r.DB.Create(&user)
	return user, nil
}

func (r *UserRepositoryImpliment) Update(user *model.User) error {
	r.DB.Save(&user)
	return nil
}

func (r *UserRepositoryImpliment) Delete(id uint64) error {
	r.DB.Where("id = ?", id).Delete(&model.User{})
	return nil
}

```

user_service.go
```go
// UserServiceの方もIUserRepository（インターフェイス）
// に依存するように変更
type UserService struct {
	UserRepository repository.IUserRepository
}

func NewUserService(repository repository.IUserRepository) UserService {
	return UserService{UserRepository: repository}
}

func (userService *UserService) GetUser(userID uint64) (model.User, error) {
	return userService.UserRepository.GetByID(userID)
}

func (UserService *UserService) CreateUser(name string) (*model.User, error) {
	user := model.User{
		Name:       name,
		Createtime: time.Now(),
		Updatetime: time.Now(),
	}
	return UserService.UserRepository.Save(&user)
}

func (UserService *UserService) UpdateUser(user *model.User) error {
	return UserService.UserRepository.Update(user)
}

func (UserService *UserService) DeleteUser(userID uint64) error {
	return UserService.UserRepository.Delete(userID)
}

```

B君
おお、難しそう。
これをすると何がどういいんですか？？

A君
アプリケーションサービスが抽象に依存しているから、もしアプリケーションサービスのテストをしたいってなったら
IUserRepositoryを実装したテスト用リポジトリ(インメモリに保存)に差し替えれば、データベースを用意しなくても
テストをすることができるよ！

B君
なるほど！

A君
他にも、データベースをRDBからNoSQLに差し替える時とかも簡単に移行できそうだね

B君
へ〜、なんとなくですがわかった気がします👍


A君
それはよかった！


具体的な実装や依存性の注入を含めたコードはこちら->https://github.com/harukitosa/ddd_sample

参考文献

ドメイン駆動設計入門　成瀬允宣
