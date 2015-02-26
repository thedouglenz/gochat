package main

type Message struct {
	Content   string
	Sender_Id string
}

func saveMessage(db DBHandle, m *Message) {
	r := db.QueryRow(`INSERT INTO "messages" ("content", "sender_id")
    VALUES ($1, $2::int) RETURNING "id"`, m.sender_id, m.content)
	var idint int64
	err := r.Scan(&idint)
	if err != nil {
		log.Print(err)
		return 0, err
	}
	return uint64(idint), err
}
