from pymongo import MongoClient


if __name__ == '__main__':
    client = MongoClient("mongodb://admin:password@localhost:27017")
    database = client.db # the new created db
    print(database)
    tables = database.list_collection_names()
        
    for count, table in enumerate(tables):
        print(f"table{count}: {table}")
    
    print("insert data in the users table")
    
    users_table = database.jobs
    
    
    mylist = [
        { "title": "DevOps Engineer", "description": "hi mom", "company": "hi dad", "salary": "1234"},
        { "title": "Cloud Engineer", "description": "hi mom", "company": "hi dad", "salary": "1234"},
        { "title": "Data Engineer", "description": "hi mom", "company": "hi dad", "salary": "1234"},
        { "title": "Container Engineer", "description": "hi mom", "company": "hi dad", "salary": "1234"},
        { "title": "Backend Engineer", "description": "hi mom", "company": "hi dad", "salary": "1234"}
    ]
    
    x = users_table.insert_many(mylist)
    print(x.inserted_ids)