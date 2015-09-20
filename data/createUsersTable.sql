CREATE TABLE 'Users' ('Name' char primary key, 'Password' char, 'Admin' integer, 'CreationDate' date);
CREATE TABLE 'Sessions' ('Name' char, 'SessionKey' char, 'LastRequest' date, primary key ('Name', 'SessionKey');
