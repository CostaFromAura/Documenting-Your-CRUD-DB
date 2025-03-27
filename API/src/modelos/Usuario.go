package modelos

import (
	"api/src/seguranca"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

// Usuario representa um usuário utilizando a rede social
type Usuario struct {
	ID       uint64    `json:"id,omitempity"`
	Nome     string    `json:"nome,omitempity`
	Nick     string    `json:"nick,omitempity`
	Email    string    `json:"email,omitempity`
	Senha    string    `json:"senha,omitempity`
	CriadoEm time.Time `json:"criadoEm,omitempity`
}

func (usuario *Usuario) Preparar(etapa string) error {
	if erro := usuario.validar(etapa); erro != nil {
		return erro
	}

	if erro := usuario.formatar(etapa); erro != nil {
		return erro
	}
	return nil
}

func (usuario *Usuario) validar(etapa string) error {
	if usuario.Nome == "" {
		return errors.New("O nome é obrigatório e não pode estar em branco")
	}
	if usuario.Nick == "" {
		return errors.New("O nick é obrigatório e não pode estar em branco")
	}
	if usuario.Email == "" {
		return errors.New("O email é obrigatório e não pode estar em branco")
	}
	if erro := checkmail.ValidateFormat(usuario.Email); erro != nil {
		return errors.New("O email inserido é invalido!")
	}

	if etapa == "cadastro" && usuario.Senha == "" {
		return errors.New("A senha é obrigatória e não pode estar em branco")
	}
	return nil
}

func (usurario *Usuario) formatar(etapa string) error {
	usurario.Nome = strings.TrimSpace(usurario.Nome)
	usurario.Nick = strings.TrimSpace(usurario.Nick)
	usurario.Email = strings.TrimSpace(usurario.Email)

	if etapa == "cadastro" {
		senhaComHash, erro := seguranca.Hash(usurario.Senha)
		if erro != nil {
			return erro
		}

		usurario.Senha = string(senhaComHash)
	}
	return nil
}
