-- nomor 1.
CREATE TABLE Data_Transaksi(
    id int PRIMARY KEY,
    Tanggal_Order datetime,
    Status_Tagihan enum('Lunas', 'Pending'),
    Tanggal_Bayar datetime
);

insert into Data_Transaksi(tanggal_order, status_tagihan, tanggal_bayar) values
     ('2022-09-26 00:00:00', 'Lunas', '2022-09-27 01:00:00'),
     ('2022-09-25 00:00:00', 'Lunas', '2022-09-26 01:00:00'),
     ('2022-10-01 00:00:00', 'Pending', null);
     
-- nomor 2.
CREATE TABLE Detail_Transaksi (
    id int PRIMARY KEY,
    ref_idTransaksi int,
    Harga int,
    Jumlah int,
    Subtotal int,
    FOREIGN KEY (ref_IdTransaksi) REFERENCES Data_Transaksi(ID)
);

INSERT INTO Detail_Transaksi(id, ref_idTransaksi, Harga, Jumlah, Subtotal) values (1, 1, 400000, 3, 1200000);

-- nomor 3.
select dt.id, td.tanggal_order, td.status, td.tanggal_pembayaran, dt.total, dt.Jumlah as jumlah_barang from Detail_Transaksi dt join Data_Transaksi td on td.id = dt.idTransaksi 
     