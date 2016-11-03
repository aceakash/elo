package elo

type Error string

const (
	PlayerAlreadyExists = Error("That name is already registered to an existing player.")
	PlayerDoesNotExist  = Error("No player by that name exists")
)

func (e Error) Error() string {
	return string(e)
}
