INSERT INTO asset_classes(id, created_at, updated_at, asset_class, description)
VALUES ('ASC_01GD16ADFNF9WEX6HSCSB2AG74', 'now', 'now', 'CASH_OR_CASH_EQUIVALENT',
        'Value of assets that are cash or can be converted to cash immediately.'),
       ('ASC_01GD16CFZ09J6GT2AHY3D8BBHF', 'now', 'now', 'COMMODITY',
        'A raw material or primary agricultural product that can be bought and sold.'),
       ('ASC_01GD16E9V3ZCDYFFPKR2DAGWJN', 'now', 'now', 'CRYPTOCURRENCY',
        'A digital or virtual currency secured by cryptography.'),
       ('ASC_01GD16EZNX0S2ERK3ZGCY5DQDZ', 'now', 'now', 'EQUITY',
        'The ownership of assets with financial value that may have debts or other liabilities attached to them.'),
       ('ASC_01GD16F5VPV8GAZ6100E0A2082', 'now', 'now', 'FIXED_INCOME',
        'Assets and securities that pay out a set level of cash flows to investors.'),
       ('ASC_01GD16FFQD6V3WEWT0540SPDHN', 'now', 'now', 'FUTURE',
        'Derivative financial contracts that obligate parties to buy or sell an asset at a predetermined future date and price.'),
       ('ASC_01GD16FMYCJKCEM7HAQ8GVMWP5', 'now', 'now', 'REAL_ESTATE', 'Property consisting of land and buildings.');

INSERT INTO auth_roles(id, created_at, updated_at, auth_role, description)
VALUES ('ATR_01GD16KFH2SCZDKN38BHRGYV8S', 'now', 'now', 'DEMO', 'A demo role with restricted rights.'),
       ('ATR_01GD16M0YMJWG32NZKC0EN5Y0Q', 'now', 'now', 'FREE', 'A free role with limited feature access.'),
       ('ATR_01GD16M7QHEE2V9Q7AQQKEAKQR', 'now', 'now', 'PLUS', 'A paying Plus role.'),
       ('ATR_01GD16MDCDAESNAYYA78R9YS63', 'now', 'now', 'PRO', 'A paying Pro role.'),
       ('ATR_01GD16MMXY0KKDFA3AVPZ81EXS', 'now', 'now', 'ENTERPRISE', 'A paying Enterprise role.'),
       ('ATR_01GD16MTAEP1611Q6W3VH9CADA', 'now', 'now', 'SUPPORT', 'A customer support role with limited permissions.'),
       ('ATR_01GD16MZBQGTGEQ5H2JNSYW1MF', 'now', 'now', 'ADMIN', 'An administrative role with elevated permissions.'),
       ('ATR_01GD16N4TC20MCKN3BA1RRBN5M', 'now', 'now', 'SUPER_ADMIN', 'A super administrative role with full rights.');

INSERT INTO transaction_types(id, created_at, updated_at, transaction_type, description)
VALUES ('TRT_01GD16SGW408D3DATDTNNX2BWD', 'now', 'now', 'BUY', 'A buy transaction.'),
       ('TRT_01GD16SQRGQ3QH4MC9WMPXJ7J9', 'now', 'now', 'SELL', 'A sell transaction.'),
       ('TRT_01GD16SWW6E2D603YZCHZZZPRV', 'now', 'now', 'STAKE', 'Staking of a cryptocurrency.'),
       ('TRT_01GD16T1VZANDKA55S02GN3K4V', 'now', 'now', 'DIVIDEND_INCOME', 'Income from dividends.'),
       ('TRT_01GD16T6S8WPD3MSXB4M7M3WHC', 'now', 'now', 'RENT_PAYMENT', 'Outflow of cash due to rent payment.'),
       ('TRT_01GD16TBXVTM18MCJ1WMWZZ00Y', 'now', 'now', 'RENT_INCOME', 'Inflow of cash due to rent collection.'),
       ('TRT_01GD16TJAVMVWP41BJZHM58GK4', 'now', 'now', 'STOCK_DIVIDEND', 'Dividends paid out in stock.');