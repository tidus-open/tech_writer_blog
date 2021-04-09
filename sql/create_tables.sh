cat /dev/null > tables.sql

echo "drop database twb_db;" >> tables.sql
echo "create database twb_db;" >> tables.sql
echo "use twb_db;" >> tables.sql

for i in {00000000..00000200} 
do	
	echo "create table if not exists twb_account_tab_$i (
        idx_user_id int unsigned auto_increment,
        user_name char(50) not null,
        passwd char(128) not null,
        delflag tinyint,
        primary key( idx_user_id),
	unique key(user_name)
        )engine=innodb, default charset=utf8;" >> tables.sql	
done	

for i in {00000000..00000001}
do
	echo "create table if not exists twb_team_tab_$i (
	idx_team_id int unsigned auto_increment,
	idx_team_name char(50) not null,
	idx_user_id int unsigned not null,
	description varchar(1000),
	create_time int unsigned not null,
	update_time int unsigned not null,
	member_count int unsigned not null,
	delflag tinyint not null,
	primary key( idx_team_id ),
	unique key (idx_user_id, idx_team_name )
	)engine=innodb, default charset=utf8;" >> tables.sql
	
done

for i in {00000000..00000200}
do
	echo "create table if not exists twb_team_member_$i (
	idx_team_user bigint unsigned,
	team_id int unsigned not null,
	user_id int unsigned not null,
	create_time int unsigned not null,
	delflag tinyint unsigned not null,
	primary key ( idx_team_user )
	)engine=innodb, default charset=utf8;" >> tables.sql
done	

for i in {00000000..00000200}
do
	echo "create table if not exists twb_member_team_$i (
	idx_user_team bigint unsigned,
	team_id int unsigned not null,
	user_id int unsigned not null,
	delflag tinyint unsigned not null,
        primary key ( idx_user_team )
	)engine=innodb, default charset=utf8;" >> tables.sql
done	

for i in {00000000..00000200}
do
	echo "create table if not exists twb_following_tab_$i (
		idx_user_following bigint unsigned,
		user_id int unsigned not null,
		following_id int unsigned not null,
		create_time int unsigned not null,
		update_time int unsigned not null,
		delflag tinyint unsigned not null,
		primary key ( idx_user_following )
	)engine=innodb, default charset=utf8;" >> tables.sql
done	

for i in {00000000..00000400}
do
	echo "create table if not exists twb_article_tab_$i (
		idx_article_id int unsigned auto_increment,
		idx_user_id int unsigned not null,
		title varchar(500) not null,
		content varchar(20000),
		create_time int unsigned not null,
		update_time int unsigned not null,
		delflag tinyint unsigned not null,
		primary key ( idx_article_id )
	)engine=innodb, default charset=utf8;" >> tables.sql
done	

for i in {00000000..00001000}
do
	echo "create table if not exists twb_comment_tab_$i (
		idx_article_create bigint unsigned,
		user_id int unsigned not null,
		create_time int unsigned not null,
		content varchar(1000) not null,
		delflag tinyint unsigned not null,
		primary key ( idx_article_create )
	)engine=innodb, default charset=utf8;" >> tables.sql
done	


for i in {00000000..00001000}
do
	echo "create table if not exists twb_article_score_tab_$i (
		idx_article_id int unsigned,
		total_score int unsigned not null,
		member_cnt int unsigned not null,
		primary key( idx_article_id )
	)engine=innodb, default charset=utf8;" >> tables.sql
done	

echo "create table if not exists twb_top_article_tab (
	idx_id int unsigned not null auto_increment,
	article_id int unsigned not null,
	article_totla_score int unsigned not null,
	article_create_time int unsigned not null,
	primary key ( idx_id )
)engine=innodb, default charset=utf8;" >> tables.sql


cat tables.sql | mysql -uentrytask -pTwb@123456
