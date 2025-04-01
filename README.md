# SW-Clone

## **Project Overview**  
This project aims to build a **SplitWise-like app** using **Microservices** architecture with a **Go** backend and a **React Native** frontend. The app will help users manage group expenses, split bills, track payments, and more. Features include user authentication, expense management, notifications, and more.

## **Features**

### **1. User Authentication (Signup, Login, OAuth Support)**  
- **Backend:** 8 Credit Points  
- **Frontend:** 7 Credit Points  
- **Description:** Users can register, log in, and authenticate via OAuth. Secure session management will be handled, including JWT-based authentication.

### **2. Create & Manage Groups (e.g., "Trip to Paris," "Roommates")**  
- **Backend:** 5 Credit Points  
- **Frontend:** 5 Credit Points  
- **Description:** Users can create and manage groups, add members, and track expenses for each group. Groups can be public or private.

### **3. Add Expenses (Title, Amount, Category, Date, Receipt Image, Individual Items, Location)**  
- **Backend:** 7 Credit Points  
- **Frontend:** 6 Credit Points  
- **Description:** Users can add expenses with details like title, amount, category, date, and an image of the receipt. Users can add individual items of the expense, and it can be tagged with the expense’s location.

### **4. Split Expenses (Equally, Custom Shares, Percentages, Divide Individual Items)**  
- **Backend:** 6 Credit Points  
- **Frontend:** 7 Credit Points  
- **Description:** Split expenses equally or customize the share percentages for each participant. Individual items can be split as well.

### **5. Expense History (Track Past Expenses with Filters and Graphs)**  
- **Backend:** 6 Credit Points  
- **Frontend:** 7 Credit Points  
- **Description:** Users can track past expenses with filters (by date, category, group, etc.). Graphs can be used for visual representation of the expenses.

### **6. Settle Balances (Mark Payments as Settled)**  
- **Backend:** 5 Credit Points  
- **Frontend:** 5 Credit Points  
- **Description:** Users can mark transactions as settled and keep track of paid and unpaid balances.

### **7. OCR for Receipt Parsing (Extract Text from Images)**  
- **Backend:** 8 Credit Points  
- **Frontend:** 6 Credit Points  
- **Description:** Users can upload receipt images, and the app will automatically extract text (e.g., amount, date, vendor) using OCR.

### **8. Expense Classification (Categorize Expenses: Breakfast, Lunch, etc.)**  
- **Backend:** 8 Credit Points  
- **Frontend:** 7 Credit Points  
- **Description:** The app automatically classifies expenses into predefined categories (e.g., "Breakfast," "Cosmetic"). Users can also create custom groups (e.g., "Food," "Entertainment").

### **9. Push Notifications (Reminders for Pending Settlements)**  
- **Backend:** 5 Credit Points  
- **Frontend:** 6 Credit Points  
- **Description:** Send push notifications for pending settlements, new expenses, or updates on group activities.

### **10. Multi-Currency Support (Convert Expenses to Default Currency)**  
- **Backend:** 7 Credit Points  
- **Frontend:** 6 Credit Points  
- **Description:** Users can view expenses in their preferred currency, with real-time currency conversion using APIs.

### **11. Export Reports (PDF/CSV Export for Tracking Expenses)**  
- **Backend:** 6 Credit Points  
- **Frontend:** 5 Credit Points  
- **Description:** Users can export their expense history in PDF or CSV formats with filters (e.g., date range, category).

### **12. Recurring Expenses (Auto-Split Monthly Bills)**  
- **Backend:** 7 Credit Points  
- **Frontend:** 6 Credit Points  
- **Description:** Users can create recurring expenses (e.g., subscriptions), and the app will automatically add them to the expense list on a regular basis (daily, weekly, monthly).

---

## **Tech Stack**

### **Backend**  
- **Language:** Go  
- **Framework:** Fiber/Echo  
- **Database:** PostgreSQL  
- **Notification Service:** Firebase Cloud Messaging or OneSignal(not confirmed)  
- **Currency Conversion API:** Open Exchange Rates or CurrencyLayer 
- **OCR:** Tesseract or other OCR APIs(not confirmed)  
- **Export Reports:** GoPDF, encoding/csv(not confirmed)

### **Frontend**  
- **Language:** JavaScript/TypeScript (React Native)  
- **State Management:** Zustand  
- **Navigation:** Expo Router  
- **Push Notifications:** Expo Push Notifications API  

### **DevOps/Containerization**  
- **Containerization:** Docker  
- **Orchestration:** Kubernetes  
- **CI/CD:** GitHub Actions  

---

## **Future Enhancements**  
- **Chat Feature:** Allow users to chat within groups to discuss expenses.  
- **Advanced AI for Expense Categorization:** Improve accuracy of automatic expense classification with machine learning models.  

---

## **Project Timeline**

| **Phase**                     | **Tasks**                                                               | **Time Estimate**      |
|-------------------------------|------------------------------------------------------------------------|------------------------|
| **1. Setup Phase**             | Backend & Frontend Setup, Dockerization, CI/CD setup                    | 2-3 weeks              |
| **2. Backend Development**     | User Authentication, Groups, Expenses, Payments, OCR, Notifications    | 5-7 weeks              |
| **3. Frontend Development**    | UI/UX Implementation, Expense Management, Notifications, Charts        | 5-7 weeks              |
| **4. Testing & Debugging**     | Backend & Frontend Testing, Integration Testing                        | 3-5 weeks              |
| **5. Deployment & DevOps**     | Kubernetes Setup, Cloud Deployment, CI/CD Integration                  | 4-6 weeks              |
| **6. Final Polish & Documentation** | UI Polish, Documentation, Launch Prep                              | 4-5 weeks              |

**Total Project Duration Estimate:** **4 - 6 months**
