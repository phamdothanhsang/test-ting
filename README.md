##### CODING TEST

### BACKEND Part

# How to run backend project
    - Navigate to main.go (press F5 , run with debug)
    - Use Postman to test API 

# What to learn to prep
    - MVVM (Vodel View Viewmodel)
    - MVC (Model View Controller)
    - Gin 
    - GORM 

# Task 
    - implement API for retrieving all Kols that are current available in the database 
    - Create dummies data based on provided Table structure (around 20-30 KOLs for the latter task)
    - Run that database on your localmachine , or any free hosting site (Use free cloud on the internet, host it for few days)
    - Return the expected response like belowed

# Expected output 
    {
    "result": "Success",
    "errorMessage": "",
    "pageIndex": 1,
    "pageSize": 100,
    "totalCount": 2,
    "KolInformation": [
        {
            "KolID": 101,
            "UserProfileID": 2001,
            "Language": "en",
            "Education": "Bachelor's in Computer Science",
            "ExpectedSalary": 70000,
            "ExpectedSalaryEnable": true,
            "ChannelSettingTypeID": 1,
            "IDFrontURL": "https://example.com/id-front.jpg",
            "IDBackURL": "https://example.com/id-back.jpg",
            "PortraitURL": "https://example.com/portrait.jpg",
            "RewardID": 301,
            "PaymentMethodID": 401,
            "TestimonialsID": 501,
            "VerificationStatus": "Verified",
            "Enabled": true,
            "ActiveDate": "2024-08-01T09:00:00Z",
            "Active": true,
            "CreatedBy": "admin",
            "CreatedDate": "2024-08-01T08:30:00Z",
            "ModifiedBy": "admin",
            "ModifiedDate": "2024-08-10T10:15:00Z",
            "IsRemove": false,
            "IsOnBoarding": true,
            "Code": "KOL2024001",
            "PortraitRightURL": "https://example.com/portrait-right.jpg",
            "PortraitLeftURL": "https://example.com/portrait-left.jpg",
            "LivenessStatus": "Passed"
        },
        {
            "KolID": 102,
            "UserProfileID": 2002,
            "Language": "vn",
            "Education": "Bachelor's in Marketing",
            "ExpectedSalary": 50000,
            "ExpectedSalaryEnable": false,
            "ChannelSettingTypeID": 2,
            "IDFrontURL": "https://example.com/id-front-2.jpg",
            "IDBackURL": "https://example.com/id-back-2.jpg",
            "PortraitURL": "https://example.com/portrait-2.jpg",
            "RewardID": 302,
            "PaymentMethodID": 402,
            "TestimonialsID": 502,
            "VerificationStatus": "Pending",
            "Enabled": true,
            "ActiveDate": "2024-08-02T09:00:00Z",
            "Active": true,
            "CreatedBy": "admin",
            "CreatedDate": "2024-08-02T08:30:00Z",
            "ModifiedBy": "admin",
            "ModifiedDate": "2024-08-11T10:15:00Z",
            "IsRemove": false,
            "IsOnBoarding": false,
            "Code": "KOL2024002",
            "PortraitRightURL": "https://example.com/portrait-right-2.jpg",
            "PortraitLeftURL": "https://example.com/portrait-left-2.jpg",
            "LivenessStatus": "Failed"
        }
        ]   
    }


### FRONTEND Part

# How to run FE project
    - Navigate to FE's main folder
    - Run this command
        - npm i
        - npm run dev
    - Navigate to localhost:3000 

# How to prep
    - Nextjs
    - Axios / fetch / ajax (anything works)


# Task (NO RESPONSIVE REQUIRED!)
    - Implement a simple component that display your response from Kols's API above
    - Implement based on this simple design
    - Two scroll button, that can scroll Kols list


