psql -dpostgres
create database credit_scoring_v2;

\c credit_scoring_v2

create table roles (
    id int not null,
    title varchar not null,
    primary key (id)
);

create table marriedStatus (
    id int not null,
    title varchar,
    primary key (id)
);

create table genderStatus (
    id int not null,
    title varchar,
    primary key (id)
);

--ภาค
create table geography (
    id int not null,
    title varchar,
    primary key (id)
);

--จังหวัด
create table provinces (
   id int generated by default as identity,
   code varchar not null,
   title varchar not null,
   title_en varchar,
   score int,
   geographyID int,
   primary key (id),
   foreign key (geographyID) references geography (id)
);
create unique index province_code_inx on provinces (code);

CREATE TABLE location (
  id int,
  subDistrict varchar not null,
  district varchar not null,
  zipcode varchar not null,
  subDistrict_code varchar not null,
  district_code varchar not null,
  province_code varchar not null,
  primary key (id),
  foreign key (province_code) references provinces (code)
);

--user
create table users (
    id bigint generated by default as identity,
    roleID int,
    --ข้อมูลเกี่ยวกับตัว user
    citizenID varchar,
    email varchar not null,
    password varchar not null,
    name varchar not null,
    surname varchar not null,
    genderID int, --รหัสเพศ fk
    marriedID int, --รหัสสถานะสมรส fk
    religion varchar, --ศาสนา
    birthday varchar not null, --วันเกิด
    phone varchar not null,
    child int, --จำนวนบุตร
    --social network
    facebook varchar,
    ig varchar,
    line varchar,--
    --ข้อมูลเกี่ยวกับที่อยู่
    address1_home varchar,
    address2_home varchar,
    subDistrict_home varchar, --เก็บชื่อตำบล
    district_home varchar, --เก็บชื่ออำเภอ
    provinceCode_home varchar, --เก็บรหัสจังหวัด
    zipCode_home varchar,
    --ข้อมูลเกี่ยวกับที่ทำงาน
    office_name varchar,
    address1_office varchar,
    address2_office varchar,
    subDistrict_office varchar,  --เก็บชื่อตำบล
    district_office varchar, --เก็บชื่ออำเภอ
    provinceCode_office varchar, --เก็บรหัสจังหวัด
    zipCode_office varchar,

    primary key (id),
    foreign key (roleID) references roles (id),
    foreign key (marriedID) references marriedStatus (id),
    foreign key (genderID) references genderStatus (id),

    foreign key (provinceCode_home) references provinces (code),

    foreign key (provinceCode_office) references provinces (code)
);
create unique index users_email_idx on users (email);
create unique index users_citizenID_idx on users (citizenID);
create unique index users_phone_idx on users (phone);

-- permission access
create table permissionAccess (
    id int,
    accessShowLoanerNewListAdmin boolean not null,
	accessShowLoanerNewListWorker boolean not null,
	accessShowLoanerInVerifyListAdmin boolean not null,
	accessShowLoanerInVerifyListWorker boolean not null,
	accessShowLoanerWaitApproveListAdmin boolean not null,
	accessShowLoanerWaitApproveListWorker boolean not null,
	verifyQuestionnaireByAdmin boolean not null,
	verifyQuestionnaireByWorker boolean not null,
	sendToApproveByAdmin boolean not null,
	sendToApproveByWorker boolean not null,
    primary key (id)
);

insert into permissionAccess (id, accessShowLoanerNewListAdmin, accessShowLoanerNewListWorker, accessShowLoanerInVerifyListAdmin, accessShowLoanerInVerifyListWorker, accessShowLoanerWaitApproveListAdmin, accessShowLoanerWaitApproveListWorker, verifyQuestionnaireByAdmin, verifyQuestionnaireByWorker, sendToApproveByAdmin, sendToApproveByWorker)
    values
        (1, true, true, true, true, true, true, true, true, true, true);

/*
    questionnaire
 */
create table  questionnaireStatus (
    id int,
    title varchar,
    primary key (id)
);

insert into questionnairestatus (id, title)
    values
        (1, 'ผู้กู้ทำแบบสอบถาม'),
        (2, 'รอการตรวจสอบข้อมูลจากพนักงาน'),
        (3, 'พนักงานกำลังตรวจสอบข้อมูล'),
        (4, 'รอการพิจารณาอนุมัติสินเชื่อ'),
        (5, 'อนุมัติ'),
        (6, 'ไม่อนุมัติ');

create index questionnaireStatus_asc on questionnairestatus (id asc);

create table criteria (
    id int generated by default as identity,
    criteria_no int not null,
    criteria_title varchar not null,
    weight int not null,
    primary key (id)
);
create unique index criteria_criteria_no_idx on criteria (criteria_no);

insert into criteria (criteria_no, criteria_title, weight)
    values
        (1, 'Character', 3),
        (2, 'Capacity', 3),
        (3, 'Capital', 2),
        (4, 'Collateral', 1),
        (5, 'Condition', 1);

create table ageOption (
    id int generated by default as identity,
    code varchar not null,
    title varchar not null,
    score int not null,
    criteria_no int not null,
    primary key (id),
    foreign key (criteria_no) references criteria (criteria_no)
);
create unique index ageOption_code_idx on ageOption (code);

insert into ageOption (code, title, score, criteria_no)
    values
        ('0', '--กรุณาระบุช่วงอายุของท่าน--', 0, 1)
        ('1', 'มากกว่า 60 ปี', 1, 1),
        ('2', '51-60 ปี', 2, 1),
        ('3', '20-30 ปี', 3, 1),
        ('4', '41-50 ปี', 4, 1),
        ('5', '31-40 ปี', 5, 1);

create table jobOption (
    id int generated by default as identity,
    code varchar not null,
    title varchar not null,
    score int,
    criteria_no int not null,
    primary key (id),
    foreign key (criteria_no) references criteria (criteria_no)
);
create unique index jobOption_code_idx on jobOption (code);

insert into jobOption (code, title, score, criteria_no)
    values
        ('0', '--กรุณาระบุอาชีพ--', 1, 1),
        ('1', 'ว่างงาน', 1, 1),
        ('2', 'นักเรียน/นิสิต/นักศึกษา', 1, 1),
        ('3', 'มีรายได้ไม่แน่นอน', 2, 1),
        ('4', 'ลูกจ้างชั่วคราว/ลูกจ้างรายวัน', 2, 1),
        ('5', 'ผู้ใช้แรงงาน/ทำงานพาร์ทไทม์', 2, 1),
        ('6', 'อาชีพอิสระ/แรงงานฝีมือ', 3, 1),
        ('7', 'เกษตรกร', 3, 1),
        ('8', 'พ่อค้าแม่ค้า', 3, 1),
        ('9', 'ข้าราชการครู', 2, 1),
        ('10', 'อาจารย์มหาวิทยาลัย', 3, 1),
        ('11', 'ข้าราชการพยาบาล', 2, 1),
        ('12', 'ตำรวจหรือทหารชั้นประทวน', 1, 1),
        ('13', 'ข้าราชการ/พนักงานราชการ', 4, 1),
        ('14', 'พนักงานประจำ/พนักงานเอกชน', 4, 1),
        ('15', 'แพทย์/สัตวแพทย์/ทันตแพทย์/เภสัชกร', 5, 1),
        ('16', 'ผู้พิพากษา/อัยการ', 5, 1),
        ('17', 'นักบินพาณิชย์', 5, 1),
        ('18', 'เจ้าของธุรกิจรายได้ต่ำกว่า 1 ล้านบาทต่อปี', 3, 1),
        ('19', 'เจ้าของธุรกิจมีรายได้ 1-10 ล้านบาทต่อปี', 4, 1),
        ('20', 'เจ้าของธุรกิจมีรายได้มากกว่า 10 ล้านบาทต่อปี', 5, 1);

create table eduOption (
    id int generated by default as identity,
    code varchar not null,
    title varchar not null,
    score int,
    criteria_no int not null,
    primary key (id),
    foreign key (criteria_no) references criteria (criteria_no)
);
create unique index eduOption_code_idx on eduOption (code);

insert into eduOption (code, title, score, criteria_no)
    values
        ('0', '--กรุณาระบุระดับการศึกษา--', 1, 1),
        ('1', 'สำเร็จการศึกษาระดับต่ำกว่า ม.6 หรือ ปวช.', 1, 1),
        ('2', 'สำเร็จการศึกษาระดับ ม.6 หรือ ปวช.', 1, 1),
        ('3', 'สำเร็จการศึกษาระดับ ปวส. / ปวท./อนุปริญญา', 1, 1),
        ('4', 'สำเร็จการศึกษาระดับปริญญาตรี', 1, 1),
        ('5', 'ระดับการศึกษาสูงกว่าปริญญาตรี', 1, 1);

create table timeJobOption (
    id int generated by default as identity,
    code varchar not null,
    title varchar not null,
    score int,
    criteria_no int not null,
    primary key (id),
    foreign key (criteria_no) references criteria (criteria_no)
);
create unique index timeJobOption_code_idx on timeJobOption (code);

insert into timeJobOption (code, title, score, criteria_no)
    values
        ('0', '--กรุณาระบุระยะเวลาที่ทำงานในอาชีพปัจจุบัน--', 1, 1),
        ('1', 'ต่ำกว่า 1 ปี', 1, 1),
        ('2', '1-2ปี', 2, 1),
        ('3', '2-4ปี', 3, 1),
        ('4', '4-6ปี', 4, 1),
        ('5', 'มากกว่า 6 ปี', 5, 1);

create table freChangeNameOption (
    id int generated by default as identity,
    code varchar not null,
    title varchar not null,
    score int,
    criteria_no int not null,
    primary key (id),
    foreign key (criteria_no) references criteria (criteria_no)
);
create unique index freChangeNameOption_code_idx on freChangeNameOption (code);

insert into freChangeNameOption (code, title, score, criteria_no)
    values
        ('0', '--จำนวนครั้งที่เคยเปลี่ยนชื่อ--', 1, 1),
        ('1', 'มากกว่า 3 ครั้ง', 1, 1),
        ('2', '3 ครั้ง', 2, 1),
        ('3', '2 ครั้ง', 3, 1),
        ('4', '1 ครั้ง', 4, 1),
        ('5', 'ไม่เคยเลย', 5, 1);

create index freChangeNameOption_code_asc on freChangeNameOption (code);

create table timeOfPhoneNumberOption (
    id int generated by default as identity,
    code varchar not null,
    title varchar not null,
    score int,
    criteria_no int not null,
    primary key (id),
    foreign key (criteria_no) references criteria (criteria_no)
);
create unique index timeOfPhoneNumberOption_code_idx on timeOfPhoneNumberOption (code);

insert into timeOfPhoneNumberOption (code, title, score, criteria_no)
    values
        (0, '--กรุณาระบุระยะเวลาที่ท่านใช้เบอร์มือถือเบอร์ปัจจุบัน--', 1, 1),
        (1, 'น้อยกว่า 1 ปี', 1, 1),
        (2, '1-2 ปี', 2, 1),
        (3, '2-5 ปี', 3, 1),
        (4, '5-10 ปี', 4, 1),
        (5, 'มากกว่า 10 ปี หรือ เบอร์ขึ้นต้น 081 หรือ 089', 5, 1);

create table timeOfNameInHouseParticularOption (
    id int generated by default as identity,
    code varchar not null,
    title varchar not null,
    score int,
    criteria_no int not null,
    primary key (id),
    foreign key (criteria_no) references criteria (criteria_no)
);
create unique index timeOfNameInHouseParticularOption_code_idx on timeOfNameInHouseParticularOption (code);

insert into timeOfNameInHouseParticularOption (code, title, score, criteria_no)
    values
        ('0', '--กรุณาระบุระยะเวลาที่ท่านอยู่อาศัยในทะเบียนบ้านปัจจุบัน--', 1, 1),
        ('1', 'ต่ำกว่า 1 ปี', 1, 1),
        ('2', '2-4 ปี', 2, 1),
        ('3', '4-6 ปี', 3, 1),
        ('4', '7-8 ปี', 4, 1),
        ('5', 'มากกว่า 8 ปี', 5, 1);

create table payDebtHistoryOption (
    id int generated by default as identity,
    code varchar not null,
    title varchar not null,
    score int,
    criteria_no int not null,
    primary key (id),
    foreign key (criteria_no) references criteria (criteria_no)
);
create unique index payDebtHistoryOption_code_idx on payDebtHistoryOption (code);

insert into payDebtHistoryOption (code, title, score, criteria_no)
    values
        ('0', '--กรุณาระบุระยะประวัติการชำระหนี้ของท่าน--', 1, 1),
        ('1', 'มีประวัติค้างชำระ,เบี้ยวหนี้,พฤติกรรมไม่โปร่งใสมากกว่า3ครั้ง/3ปี', 1, 1),
        ('2', 'มีประวัติค้างชำระ,เบี้ยวหนี้,พฤติกรรมไม่โปร่งใส 3ครั้ง/3ปี', 2, 1),
        ('3', 'มีประวัติค้างชำระ,เบี้ยวหนี้,พฤติกรรมไม่โปร่งใส 2ครั้ง/3ปี', 3, 1),
        ('4', 'มีประวัติค้างชำระ,เบี้ยวหนี้,พฤติกรรมไม่โปร่งใส 1ครั้ง/3ปี', 4, 1),
        ('5', 'ไม่มีประวัติค้างชำระ,เบี้ยวหนี้,พฤติกรรมไม่โปร่งใส ภายใน 3ปี', 5, 1);

create table statusInHouseParticularOption (
    id int generated by default as identity,
    code varchar not null,
    title varchar not null,
    score int,
    criteria_no int not null,
    primary key (id),
    foreign key (criteria_no) references criteria (criteria_no)
);
create unique index statusInHouseParticularOption_code_idx on statusInHouseParticularOption (code);

insert into statusInHouseParticularOption (code, title, score, criteria_no)
    values
        ('0', '--กรุณาระบุข้อมูลการค้ำประกันให้ผู้อื่นของท่าน--', 1, 1),
        ('1', 'ไม่มีชื่ออยู่ในทะเบียนบ้านและไม่ใช่เจ้าของที่ดิน', 1, 1),
        ('2', 'เป็นผู้อยู่อาศัยและไม่ใช่เจ้าของที่ดิน', 2, 1),
        ('3', 'เป็นเจ้าบ้านและไม่ใช่เจ้าของที่ดิน หรือ เป็นเจ้าของกรรมสิทธิ์ในที่อยู่อาศัยเช่น คอนโดมิเนี่ยม', 3, 1),
        ('4', 'เป็นผู้อยู่อาศัยในทะเบียนบ้านและเป็นเจ้าของที่ดิน', 4, 1),
        ('5', 'เป็นเจ้าบ้านและเจ้าของที่ดิน', 5, 1);

 create table incomePerDebtOption (
    id int generated by default as identity,
    code varchar not null,
    title varchar not null,
    score int,
    criteria_no int not null,
    primary key (id),
    foreign key (criteria_no) references criteria (criteria_no)
);
create unique index incomePerDebtOption_code_idx on incomePerDebtOption (code);

insert into incomePerDebtOption (code, title, score, criteria_no)
    values
        ('1', 'รายได้สุทธิต่อเดือน < หรือ = จำนวนเงินผ่อนชำระต่อเดือน', 1, 2),
        ('2', 'รายได้สุทธิต่อเดือน > จำนวนเงินผ่อนชำระ = 1-20%', 2, 2),
        ('3', 'รายได้สุทธิต่อเดือน>จำนวนเงินผ่อนชำระ อยู่ในช่วง 21 -50%', 3, 2),
        ('4', 'รายได้สุทธิต่อเดือน>จำนวนเงินผ่อนชำระ อยู่ในช่วง 51-80%', 4, 2),
        ('5', 'รายได้ประจำสุทธิต่อเดือน>จำนวนเงินผ่อนชำระ = มากกว่า 80%ขึ้นไป', 5, 2);

create table totalDebtPerYearIncomeOption (
    id int generated by default as identity,
    code varchar not null,
    title varchar not null,
    score int,
    criteria_no int not null,
    primary key (id),
    foreign key (criteria_no) references criteria (criteria_no)
);
create unique index totalDebtPerYearIncomeOption_code_idx on totalDebtPerYearIncomeOption (code);

insert into totalDebtPerYearIncomeOption (code, title, score, criteria_no)
    values
        ('1', 'ยอดผ่อนชำระหนี้สินต่างๆ> 40 % ของรายได้', 1, 2),
        ('2', 'ยอดผ่อนชำระหนี้สินต่างๆ= 40 % ของรายได้', 2, 2),
        ('3', 'ยอดผ่อนชำระหนี้สินต่างๆ < 40 % ของรายได้', 3, 2),
        ('4', 'ยอดผ่อนชำระหนี้สินต่างๆ < 30-39 % ของรายได้', 4, 2),
        ('5', 'ยอดผ่อนชำระหนี้สินต่างๆ < 0-29 % ของรายได้', 5, 2);

create table savingPerLoanOption (
    id int generated by default as identity,
    code varchar not null,
    title varchar not null,
    score int,
    criteria_no int not null,
    primary key (id),
    foreign key (criteria_no) references criteria (criteria_no)
);
create unique index savingPerLoanOption_code_idx on savingPerLoanOption (code);

insert into savingPerLoanOption (code, title, score, criteria_no)
    values
        ('1', 'ไม่มีเงินเก็บออมหรือทรัพย์สินปลอดภาระหนี้สำรองไว้ยามฉุกเฉิน', 1, 3),
        ('2', 'มีเงินออมหรือทรัพย์สินปลอดภาระหนี้ไว้อย่างน้อย 5-9% ของวงเงินกู้ ', 2, 3),
        ('3', 'มีเงินออมหรือทรัพย์สินปลอดภาระหนี้ไว้ 10-14% ของวงเงินกู้ ', 3, 3),
        ('4', 'เงินออมหรือทรัพย์สินปลอดภาระหนี้ไว้ 15-20% ของวงเงินกู้', 4, 3),
        ('5', 'มีเงินออมหรือทรัพย์สินปลอดภาระหนี้มากกว่า 20% ของวงเงินกู้ ', 5, 3);

create table mortgageSecuritiesPerLoanOption (
    id int generated by default as identity,
    code varchar not null,
    title varchar not null,
    score int,
    criteria_no int not null,
    primary key (id),
    foreign key (criteria_no) references criteria (criteria_no)
);
create unique index mortgageSecuritiesPerLoanOption_code_idx on mortgageSecuritiesPerLoanOption (code);

insert into mortgageSecuritiesPerLoanOption (code, title, score, criteria_no)
    values
        ('1', 'มูลค่าหลักทรัพย์ค้ำประกันน้อยกว่าวงเงินกู้ หรือไม่มีหลักทรัพย์ค้ำประกัน', 1, 4 ),
        ('4', 'มูลค่าหลักทรัพย์ค้ำประกัน = วงเงินกู้', 4, 4 ),
        ('5', 'มูลค่าหลักทรัพย์ค้ำประกัน > วงเงินกู้', 5, 4 );

create table haveGuarantorOption (
    id int generated by default as identity,
    code varchar not null,
    title varchar not null,
    score int,
    criteria_no int not null,
    primary key (id),
    foreign key (criteria_no) references criteria (criteria_no)
);
create unique index haveGuarantorOption_code_idx on haveGuarantorOption (code);

insert into haveGuarantorOption (code, title, score, criteria_no)
    values
        ('0', '--ท่านมีบุคคลหรือสถาบันการเงินค้ำประกันหรือไม่--', 1, 4),
        ('1', 'ไม่มีบุคคลหรือสถาบันการเงินค้ำประกัน', 1, 4),
        ('5', 'มีบุคคลที่น่าเชื่อถือค้ำประกัน', 5, 4);

create table iamGuarantorOption (
    id int generated by default as identity,
    code varchar not null,
    title varchar not null,
    score int,
    criteria_no int not null,
    primary key (id),
    foreign key (criteria_no) references criteria (criteria_no)
);
create unique index iamGuarantorOption_code_idx on iamGuarantorOption (code);

insert into iamGuarantorOption (code, title, score, criteria_no)
    values
        ('0', '--กรุณาระบุข้อมูลการค้ำประกันให้ผู้อื่นของท่าน--', 1, 4),
        ('3', 'ท่านมีการค้ำประกันให้บุคคลอื่น', 3, 4),
        ('5', 'ท่านไม่ได้ค้ำประกันให้บุคคลอื่น', 5, 4);

create table incomeTrendOption (
    id int generated by default as identity,
    code varchar not null,
    title varchar not null,
    score int,
    criteria_no int not null,
    primary key (id),
    foreign key (criteria_no) references criteria (criteria_no)
);
create unique index incomeTrendOption_code_idx on incomeTrendOption (code);

insert into incomeTrendOption (code, title, score, criteria_no)
    values
        ('0', '--กรุณาประเมินแนวโน้มเศรษฐกิจที่มีผลต่อรายได้ของท่าน--', 1, 5),
        ('1', 'รายได้ลดลงมากกว่า 10%/ปี', 1, 5),
        ('2', 'รายได้ลดลง 1-9%/ปี', 2, 5),
        ('3', 'รายได้คงเดิม', 3, 5),
        ('4', 'รายได้เพิ่มขึ้น 1-9%', 4, 5),
        ('5', 'รายได้เพิ่มขึ้น 10%ขึ้นไป/ปี', 5, 5);

create table loanObjectOption (
    id int generated by default as identity,
    code varchar not null,
    title varchar not null,
    score int,
    criteria_no int not null,
    primary key (id),
    foreign key (criteria_no) references criteria (criteria_no)
);
create unique index loanObjectOption_code_idx on loanObjectOption (code);

insert into loanObjectOption (code, title, score, criteria_no)
    values
        ('0', '--กรุณาระบุวัตถุประสงค์ในการกู้ยืมของท่าน--', 1, 5),
        ('1', 'ซื้อสิ่งของฟุ่มเพือยหรือเครื่องประดับ, ท่องเที่ยว, การศึกษา, ลงทุนใน ,Forex, option, หุ้น, เงินดิจิตอล', 1, 5),
        ('2', 'เพื่ออุปโภคหรือการบริโภค/ลงทุนใดๆเพื่อการสร้างอาชีพใหม่', 2, 5),
        ('3', 'เพื่ออุปโภคหรือการบริโภค/ลงทุนใดๆเพื่อการสร้างอาชีพใหม่', 3, 5),
        ('4', 'เพื่อเสริมสภาพคล่องในการดำเนินธุรกิจหรือกิจการส่วนตัว', 4, 5),
        ('5', 'เพื่อขยายกิจการหรือธุรกิจ/ เพื่อลงทุนในการเพิ่มศักยภาพการแข่งขันในเชิงธุรกิจ', 5, 5);

create table questionnaire (
    id bigint generated by default as identity,
    loanerID bigint not null, --ผู้กู้
    updatedBy bigint, --ผู้ตรวจสอบล่าสุด
    approveBy bigint, --ผู้อนุมัติ

    statusID int not null,
    sendAt timestamp with time zone,
    updatedAt timestamp with time zone,

    approveRate float, --อนุมัติเปอร์เซนต์
    approveTotal float, --อนุมัติทั้งหมด (บาท)
    interest float, --อัตราดอกเบี้ย
    loanerPayback float, --จำนวนเงินที่ผู้กู้ต้องชำระคืน

    verifyComment varchar, --หมายเหตุการตรวจสอบ
    approveComment varchar, --หมายเหตุการอนุมัติ

    creditGrade    varchar, --[A,B,C,D,F]
    creditRisk  varchar, --[1,2,3,4,5]
    riskLevel   varchar, --[High, Medium ...]
    matrixIndex  varchar, --[A1, A2, ... F5]

    -- ส่วนผู้กู้กรอก
    --
    suggest varchar not null default '', --ผู้แนะนำ
    suggestW varchar not null default '', --ผู้แนะนำ
    suggestScore int, --คะแนนของผู้แนะนำ
    suggestGiveScore int, --คะแนนที่ผู้แนะนำให้กับ user

    -- ข้อมูลตัวเลขกรอกมือ
    income float not null,
    loan float not null,
    debtPerMonth float not null,
    totalDebt float not null,
    saving float not null,
    mortgageSecurities float not null,

    -- ข้อมูลตัวเลือก (option)
    age varchar not null,
    job varchar not null,
    edu varchar not null,
    timeJob varchar not null,
    freChangeName varchar not null,
    timeOfPhoneNumber varchar not null,
    timeOfNameInHouseParticular varchar not null,
    payDebtHistory varchar not null,
    statusInHouseParticular varchar not null,

    incomePerDebt varchar not null,
    totalDebtPerYearIncome varchar not null,
    savingPerLoan varchar not null,
    mortgageSecuritiesPerLoan varchar not null,

    haveGuarantor varchar not null,
    iamGuarantor varchar not null,
    incomeTrend varchar not null,
    loanObject varchar not null,
    provinceCode varchar not null,

    -- ส่วนผู้ตรวจสอบแก้ไข
    incomeW float not null,
    loanW float not null,
    debtPerMonthW float not null,
    totalDebtW float not null,
    savingW float not null,
    mortgageSecuritiesW float not null,

    -- ข้อมูลตัวเลือก (option)
    ageW varchar not null,
    jobW varchar not null,
    eduW varchar not null,
    timeJobW varchar not null,
    freChangeNameW varchar not null,
    timeOfPhoneNumberW varchar not null,
    timeOfNameInHouseParticularW varchar not null,
    payDebtHistoryW varchar not null,
    statusInHouseParticularW varchar not null,

    incomePerDebtW varchar not null,
    totalDebtPerYearIncomeW varchar not null,
    savingPerLoanW varchar not null,
    mortgageSecuritiesPerLoanW varchar not null,

    haveGuarantorW varchar not null,
    iamGuarantorW varchar not null,
    incomeTrendW varchar not null,
    loanObjectW varchar not null,
    provinceCodeW varchar not null,

    primary key (id),
    foreign key (loanerID) references users (id),
    foreign key (provinceCode) references provinces (code),
    foreign key (statusID) references questionnaireStatus (id),


    foreign key (age) references ageOption (code),
    foreign key (job) references jobOption (code),
    foreign key (edu) references eduOption (code),
    foreign key (timeJob) references timeJobOption (code),
    foreign key (freChangeName) references freChangeNameOption (code),
    foreign key (timeOfPhoneNumber) references timeOfPhoneNumberOption (code),
    foreign key (timeOfNameInHouseParticular) references timeOfNameInHouseParticularOption (code),
    foreign key (payDebtHistory) references payDebtHistoryOption (code),
    foreign key (statusInHouseParticular) references statusInHouseParticularOption (code),
    foreign key (incomePerDebt) references incomePerDebtOption (code),
    foreign key (totalDebtPerYearIncome) references totalDebtPerYearIncomeOption (code),
    foreign key (savingPerLoan) references savingPerLoanOption (code),
    foreign key (mortgageSecuritiesPerLoan) references mortgageSecuritiesPerLoanOption (code),
    foreign key (haveGuarantor) references haveGuarantorOption (code),
    foreign key (iamGuarantor) references iamGuarantorOption (code),
    foreign key (incomeTrend) references incomeTrendOption (code),
    foreign key (loanObject) references loanObjectOption (code),
    foreign key (provinceCode) references provinces (code),


    foreign key (ageW) references ageOption (code),
    foreign key (jobW) references jobOption (code),
    foreign key (eduW) references eduOption (code),
    foreign key (timeJobW) references timeJobOption (code),
    foreign key (freChangeNameW) references freChangeNameOption (code),
    foreign key (timeOfPhoneNumberW) references timeOfPhoneNumberOption (code),
    foreign key (timeOfNameInHouseParticularW) references timeOfNameInHouseParticularOption (code),
    foreign key (payDebtHistoryW) references payDebtHistoryOption (code),
    foreign key (statusInHouseParticularW) references statusInHouseParticularOption (code),
    foreign key (incomePerDebtW) references incomePerDebtOption (code),
    foreign key (totalDebtPerYearIncomeW) references totalDebtPerYearIncomeOption (code),
    foreign key (savingPerLoanW) references savingPerLoanOption (code),
    foreign key (mortgageSecuritiesPerLoanW) references mortgageSecuritiesPerLoanOption (code),
    foreign key (haveGuarantorW) references haveGuarantorOption (code),
    foreign key (iamGuarantorW) references iamGuarantorOption (code),
    foreign key (incomeTrendW) references incomeTrendOption (code),
    foreign key (loanObjectW) references loanObjectOption (code),
    foreign key (provinceCodeW) references provinces (code)
);
create unique index questionnaire_loanerID_idx on questionnaire (loanerID);

create table funcAllow (
    id int generated by default as identity,
    funcNo int not null,
    title varchar not null,
    superAdminAllow bool default true,
    adminAllow bool default true,
    employeeAllow bool default true,
    loanerAllow bool default true,
    primary key (id)
);
create unique index funcAllow_funcNo_idx on funcAllow (funcNo);






