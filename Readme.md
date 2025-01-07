# PT XYZ Multifinance Scenario Test

## **Overview**
This is a Go-based backend application for paylater feature

## **Features**
- See credit limit
- Do transactional with credit limit
- Edit profile

## **Technologies Used**
- Go
- MySQL
- Swagger/OpenAPI for documentation
- Testify for unit testing

## **Setup Instructions**
1. Clone the repository:
   ```bash
   git clone https://gitlab.com/fabrianivan/xyz-multifinance.git
   ```
   or
   ```bash
   git clone https://github.com/fabrianivan-id/xyz-multifinance.git
   ```

## **Architecture of Application**

+-------------------+
|   Client (Web)    |
|   (Frontend)      |
+-------------------+
          |
          v
+-------------------+
|   API Gateway     |
|   (HTTP Server)   |
+-------------------+
          |
          v
+-------------------+
|   Service Layer   |
|  (Business Logic) |
|                   |
|  +--------------+ |
|  | Consumer     | |
|  | Service      | |
|  +--------------+ |
|  +--------------+ |
|  | Transaction  | |
|  | Service      | |
|  +--------------+ |
+-------------------+
          |
          v
+-------------------+
|   Repository Layer|
|   (Data Access)   |
|                   |
|  +--------------+ |
|  | Consumer     | |
|  | Repository   | |
|  +--------------+ |
|  +--------------+ |
|  | Transaction  | |
|  | Repository   | |
|  +--------------+ |
+-------------------+
          |
          v
+-------------------+
|   Database Layer  |
|   (PostgreSQL)    |
+-------------------+

## **Database Structure**

+-------------------+
|     Consumers     |
+-------------------+
| NIK (PK)          |
| Full_Name         |
| Legal_Name        |
| Place_Of_Birth    |
| Date_Of_Birth     |
| Salary            |
| Limit_1           |
| Limit_2           |
| Limit_3           |
| Limit_4           |
+-------------------+
          |
          | 1
          |
          | N
+-------------------+
|   Transactions     |
+-------------------+
| Contract_Number (PK) |
| NIK (FK)          |
| OTR               |
| Admin_Fee         |
| Installment       |
| Interest          |
| Asset_Name        |
+-------------------+