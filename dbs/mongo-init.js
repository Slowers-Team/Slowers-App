db.createUser({
  user: 'the_username',
  pwd: 'the_password',
  roles: [
    {
      role: 'dbOwner',
      db: 'default_db',
    },
    {
      role: 'dbOwner',
      db: 'Slowers',
    },
    {
      role: 'dbOwner',
      db: 'SlowersTest',
    },
  ],
});
