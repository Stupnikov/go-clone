{
	// Open the database.
	db, err := bolt.Open(tempfile(), 0666, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer os.Remove(db.Path())

	// Start a read-write transaction.
	if err := db.Update(func(tx *bolt.Tx) error {
		// Create a new bucket.
		b, err := tx.CreateBucket([]byte("animals"))
		if err != nil {
			return err
		}

		// Insert data into a bucket.
		if err := b.Put([]byte("dog"), []byte("fun")); err != nil {
			log.Fatal(err)
		}
		if err := b.Put([]byte("cat"), []byte("lame")); err != nil {
			log.Fatal(err)
		}
		if err := b.Put([]byte("liger"), []byte("awesome")); err != nil {
			log.Fatal(err)
		}

		// Create a cursor for iteration.
		c := b.Cursor()

		// Iterate over items in sorted key order. This starts from the
		// first key/value pair and updates the k/v variables to the
		// next key/value on each iteration.
		//
		// The loop finishes at the end of the cursor when a nil key is returned.
		for k, v := c.First(); k != nil; k, v = c.Next() {
			fmt.Printf("A %s is %s.\n", k, v)
		}

		return nil
	}); err != nil {
		log.Fatal(err)
	}

	if err := db.Close(); err != nil {
		log.Fatal(err)
	}

	// Output:
	// A cat is lame.
	// A dog is fun.
	// A liger is awesome.
}