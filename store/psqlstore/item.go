package psqlstore

import (
	"context"
	"log"

	"github.com/djermy/runeprofitforge-golang-api/runescape"
)

func (s *PsqlStore) GetItems() ([]runescape.Item, error) {
	var items []runescape.Item

	rows, err := s.Conn.Query(
		context.Background(),
		`
		SELECT
			id,
			name,
			type,
			members,
			description,
			icon,
			icon_large,
			type_icon,
			current,
			today
		FROM
			item
		;`,
	)

	if err != nil {
		log.Println(err)
		return []runescape.Item{}, err
	}
	defer rows.Close()

	for rows.Next() {
		var item runescape.Item
		err := rows.Scan(
			&item.ID,
			&item.Name,
			&item.Type,
			&item.Members,
			&item.Description,
			&item.Icon,
			&item.IconLarge,
			&item.TypeIcon,
			&item.Current,
			&item.Today,
		)
		if err != nil {
			log.Println(err)
			return []runescape.Item{}, err
		}

		items = append(items, item)
	}
	return items, err
}

func (s *PsqlStore) GetItem(id string) (runescape.Item, error) {
	var item runescape.Item
	rows, err := s.Conn.Query(
		context.Background(),
		`
		SELECT
			id,
			name,
			type,
			members,
			description,
			icon,
			icon_large,
			type_icon,
			current,
			today
		FROM
			item
		WHERE
			id = $1
		LIMIT 1
		;`,
		id,
	)

	if err != nil {
		log.Println(err)
		return runescape.Item{}, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(
			&item.ID,
			&item.Name,
			&item.Type,
			&item.Members,
			&item.Description,
			&item.Icon,
			&item.IconLarge,
			&item.TypeIcon,
			&item.Current,
			&item.Today,
		)
		if err != nil {
			log.Println(err)
			return runescape.Item{}, err
		}
	}

	return item, nil
}

func (s *PsqlStore) CreateItem(item *runescape.Item) error {
	err := s.Conn.QueryRow(
		context.Background(),
		`
		INSERT INTO item (
			id,
			name,
			type,
			members,
			description,
			icon,
			icon_large,
			type_icon,
			current,
			today
		) VALUES (
			$1,
			$2,
			$3,
			$4,
			$5,
			$6,
			$7,
			$8,
			$9,
			$10
		)
		RETURNING id
		;`,
		item.ID,
		item.Name,
		item.Type,
		item.Members,
		item.Description,
		item.Icon,
		item.IconLarge,
		item.TypeIcon,
		item.Current,
		item.Today,
	).Scan(&item.ID)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (s *PsqlStore) UpdateItem(id string, item *runescape.Item) error {
	rows, err := s.Conn.Query(
		context.Background(),
		`
		UPDATE item
		SET 
			name = $2,
			type = $3,
			members = $4,
			description = $5,
			icon = $6,
			icon_large = $7,
			type_icon = $8,
			current = $9,
			today = $10
		WHERE
			id = $1;
			`,
		item.ID,
		item.Name,
		item.Type,
		item.Members,
		item.Description,
		item.Icon,
		item.IconLarge,
		item.TypeIcon,
		item.Current,
		item.Today,
	)
	if err != nil {
		log.Println(err)
		return err
	}
	defer rows.Close()

	return nil
}

func (s *PsqlStore) DeleteItem(id string) error {
	rows, err := s.Conn.Query(
		context.Background(),
		`
		DELETE FROM item
		WHERE
			id = $1;
			`,
		id,
	)
	if err != nil {
		log.Println(err)
		return err
	}
	defer rows.Close()

	return nil
}
