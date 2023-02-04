CREATE TABLE usersinfo (
    uid  SERIAL PRIMARY KEY,
    ufname VARCHAR(50),
    ulname VARCHAR(50)
    )


/*
IF  NOT EXISTS (SELECT * FROM sys.objects 
WHERE object_id = OBJECT_ID(N'[usersInfoDB].[usersinfo]') AND type in (N'U'))

BEGIN
CREATE TABLE [usersInfoDB].[usersinfo](
    uid  SEQUENCE PRIMARY KEY,
    ufname VARCHAR(50),
    ulname VARCHAR(50),
) 

END
*/


/*
IF object_id('usersinfo', 'U') is not null
    PRINT 'The table usersinfo already exists'
ELSE
    CREATE TABLE usersinfo (
    uid  INT PRIMARY KEY,
    ufname VARCHAR(50),
    ulname VARCHAR(50)
    )
*/
