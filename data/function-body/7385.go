{
	db := MustOpenDB()
	defer db.MustClose()
	if err := db.Update(func(tx *bolt.Tx) error {
		if _, err := tx.CreateBucketIfNotExists([]byte{}); err != bolt.ErrBucketNameRequired {
			t.Fatalf("unexpected error: %s", err)
		}

		if _, err := tx.CreateBucketIfNotExists(nil); err != bolt.ErrBucketNameRequired {
			t.Fatalf("unexpected error: %s", err)
		}

		return nil
	}); err != nil {
		t.Fatal(err)
	}
}