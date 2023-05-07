db = db.getSiblingDB('sample_db')

db.createUser({
    user: 'some_user',
    pwd: 'random_pass',
    roles: [
      {
        role: 'dbOwner',
      db: 'sample_db',
    },
  ],
});