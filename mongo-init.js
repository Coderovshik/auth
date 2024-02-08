db.createUser(
    {
        user: "foo",
        pwd: "bar",
        roles: [
            {
                role: "readWrite",
                db: "auth"
            }
        ]
    }
);