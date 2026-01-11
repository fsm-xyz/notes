### SQLlite的基础内容和面试点 (SQLite Basics and Interview Points)

#### SQLlite基础内容 (SQLite Basics)

SQLite是一个轻量级的、无服务器的、自给自足的、零配置的事务性SQL数据库引擎。它不需要单独的服务器进程，可以在应用程序中直接使用。以下是关于SQLite的一些核心基础知识：

**1. 架构 (Architecture):**
* **无服务器 (Serverless):** 与MySQL、PostgreSQL不同，SQLite不需要一个独立的数据库服务器进程。数据库引擎作为库直接嵌入到应用程序中。
* **零配置 (Zero-Configuration):** 不需要进行复杂的安装或配置。创建一个数据库文件（`.db`或`.sqlite`）就可以开始使用。
* **事务性 (Transactional):** 支持完整的ACID（原子性、一致性、隔离性、持久性）事务。

**2. 数据类型 (Data Types):**
SQLite使用更灵活的动态类型系统，称为**Manifest Typing**。一个值的类型取决于它被存储的方式，而不是它的列的数据类型。SQLite支持以下五种主要数据类型：
* **NULL:** 值为空。
* **INTEGER:** 带符号整数。
* **REAL:** 浮点数。
* **TEXT:** 文本字符串。
* **BLOB:** 二进制大对象，用于存储原始数据。

**3. 基本SQL操作 (Basic SQL Operations):**
SQLite支持标准的SQL语法，包括：
* **DDL (Data Definition Language):**
    * `CREATE TABLE`: 创建表。
    * `ALTER TABLE`: 修改表结构。
    * `DROP TABLE`: 删除表。
* **DML (Data Manipulation Language):**
    * `INSERT INTO`: 插入数据。
    * `SELECT`: 查询数据。
    * `UPDATE`: 更新数据。
    * `DELETE FROM`: 删除数据。
* **事务控制 (Transaction Control):**
    * `BEGIN TRANSACTION`: 开始一个事务。
    * `COMMIT`: 提交事务，使更改永久化。
    * `ROLLBACK`: 回滚事务，撤销所有更改。

**4. 数据库文件 (Database File):**
整个数据库（包括所有表、索引和视图）都存储在一个单一的磁盘文件中。这使得数据库的管理、备份和传输变得非常简单。

---

### SQLlite面试点 (SQLite Interview Points)

面试官可能会从多个角度考察你对SQLite的理解。除了基础概念，以下是一些常见的面试问题和考点：

**1. 核心概念对比 (Core Concept Comparison):**
* **SQLite与MySQL/PostgreSQL的区别是什么？**
    * **SQLite:** 无服务器、嵌入式、单文件、零配置、适用于小型应用和设备。
    * **MySQL/PostgreSQL:** 客户端-服务器架构、需要独立的服务器、适用于大型、高并发的Web应用。
* **SQLite的优缺点是什么？**
    * **优点:** 易于使用、部署简单、占用资源少、高可靠性、适用于移动应用、桌面应用和物联网设备。
    * **缺点:** 不适合高并发（只支持一个写操作）、缺乏用户管理和权限控制、性能在大数据量下不如客户端-服务器数据库。

**2. 数据类型 (Data Types):**
* **解释一下SQLite的Manifest Typing？**
    * SQLite的类型系统更像是一种“建议”而不是强制。你可以将一个字符串存储在声明为`INTEGER`的列中，SQLite会尝试转换，如果转换失败，它仍然会将其作为文本存储。
    * 比如，`CREATE TABLE test(a INTEGER);` 后，`INSERT INTO test VALUES('abc');` 是可以成功的。

**3. 并发控制 (Concurrency Control):**
* **SQLite如何处理并发写入？**
    * SQLite在任何时刻只允许一个进程进行写操作。当一个进程开始写操作时，它会锁定整个数据库文件。其他写操作必须等待。
    * **文件锁定 (File Locking):** SQLite使用操作系统提供的文件锁定机制来控制并发。
    * 这个问题通常用来考察你对SQLite“非高并发”特性的理解。

**4. 事务与ACID (Transactions and ACID):**
* **SQLite是如何保证ACID的？**
    * **原子性 (Atomicity):** 使用回滚日志（Rollback Journal）或预写日志（WAL）来确保事务要么全部成功，要么全部失败。
    * **一致性 (Consistency):** 依赖于原子性和隔离性来保证数据从一个有效状态转移到另一个有效状态。
    * **隔离性 (Isolation):** 通过锁定机制，确保一个事务的中间状态不会被其他并发事务看到。
    * **持久性 (Durability):** 提交后，数据更改会写入磁盘，即使系统崩溃也不会丢失。

**5. 性能优化 (Performance Optimization):**
* **如何优化SQLite查询？**
    * **使用索引 (Index):** 在`WHERE`子句中经常用到的列上创建索引。
    * **理解`EXPLAIN QUERY PLAN`:** 使用这个命令可以查看SQLite如何执行你的查询，从而帮助你发现性能瓶颈。
    * **避免全表扫描 (Full Table Scans):** 遵循SQL最佳实践，尽量通过索引进行查询。
    * **考虑使用`WAL` (Write-Ahead Logging)模式:** `WAL`模式可以显著提高并发读取的性能，因为它允许读操作和写操作同时进行，读操作不会阻塞写操作。

---