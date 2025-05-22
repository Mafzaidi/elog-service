package account

type UseCase interface {
	Store(pl *CreateParams) error
}
