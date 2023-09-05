package main

import "context"

func main() {
	db, err := Connect()
	if err != nil {
		panic(err)
	}

	// membuat konteks kosong
	ctx := context.Background()

	// memulai transaksi menggunakan method `BeginTx`
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		panic(err)
	}

	// panggil method `Rollback` jika terjadi error
	defer tx.Rollback()

	// step 1: mendapatkan data produk yang ingin dibeli
	rows, err = tx.QueryContext(ctx, "SELECT * FROM products WHERE id = 1")
	if err != nil {
		// tx.Rollback() // jika terjadi error, rollback transaksi bisa menggunakan ini disetiap error yang terjadi atau dengan cara defer
		panic(err)
	}

	// step 2: mengurangi/update stok produk yang dibeli
	_, err = tx.ExecContext(ctx, "UPDATE products SET stock = stock - 1 WHERE id = 1")
	if err != nil {
		panic(err)
	}

	// step 3: menambahkan data ke tabel sales / transaksi penjualan
	_, err = tx.ExecContext(ctx, "INSERT INTO sales (product_id, quantity, total_price) VALUES (1, 1, 10000)")
	if err != nil {
		panic(err)
	}

	// step 4: menambahkan data ke tabel transactions / transaksi pembayaran (pembelian) produk
	_, err = tx.ExecContext(ctx, "INSERT INTO transactions (date, product_id, amount) VALUES ('2020-01-01', 1, 15000)")
	if err != nil {
		panic(err)
	}

	// menutup transaksi dengan method `Commit`
	err = tx.Commit()
	if err != nil {
		panic(err)
	}
}
