DROP TABLE IF EXISTS pets;
CREATE TABLE pets(
    name VARCHAR(255) NOT NULL,
    specie VARCHAR(127) NOT NULL,
    sex VARCHAR(15) NOT NULL,
    birthdate TIMESTAMP NOT NULL,

    id VARCHAR(31) NOT NULL UNIQUE,

    PRIMARY KEY(id)
);

INSERT INTO pets (name, specie, sex, birthdate, id) VALUES
('mascota01','lechat','male','2022-09-20T00:00:00Z','2F587Vez33hOY3J4UnmOjspA2jy'),
('mascota02','lechat','male','2022-09-19T00:00:00Z','2F58Etj1ItiHPmBSDZZpTmER6dc'),
('mascota03','lechat','female','2022-09-20T00:00:00Z','2F58EvkHL18YPf3Mnp1JYF7Grq4'),
('mascota04','lechat','male','2022-09-18T00:00:00Z','2F58F9eUqEOUUMUw6VysVjdRWTC'),
('mascota05','lechat','male','2022-09-17T00:00:00Z','2F58F6AZrmazs3m5GAvcPXkBtFT'),
('mascota06','dogge','female','2022-09-10T00:00:00Z','2F58FHJDbI0LEFlZ6RwXOSufnfT'),
('mascota07','dogge','female','2022-09-09T00:00:00Z','2F58FNdcKwEcB2IAPKcoGwAAGCv'),
('mascota08','dogge','female','2022-09-09T00:00:00Z','2F58FOBnnXjRcaQP09nbfkXEyLW'),
('mascota09','dogge','male','2022-09-09T00:00:00Z','2F58FVMJw5I6uinhoxegkGCcVxH'),
('mascota10','iguanatango','female','2022-09-08T00:00:00Z','2F58FU5Shfp78OG3nh89lqM8fj2'),
('mascota11','iguanatango','male','2022-08-20T00:00:00Z','2F58FapsMoU78O2Y68hjdv5udDT'),
('mascota12','iguanatango','male','2022-08-20T00:00:00Z','2F58FhUnIDJKWpyO32x8T4Db1Lh'),
('mascota13','birdo','male','2022-08-20T00:00:00Z','2F58FgHnMKQ9erWbnFYNozgfrdy'),
('mascota14','birdo','female','2022-07-20T00:00:00Z','2F58FtkVy5aLkcHwGHqFKTwhIcK'),
('mascota15','alpaca','female','2022-09-15T00:00:00Z','2F58FvldIAZDK0SrceavyGgVfYz'),
('mascota16','crabbo','male','2022-09-15T00:00:00Z','2F58Fz97EEJh1ySW2FnaqlWteu6'),
('mascota17','birdo','female','2022-09-15T00:00:00Z','2F58GgWANNlpoCJCfQkbm8N4siQ');