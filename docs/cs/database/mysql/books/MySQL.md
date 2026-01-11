# MySQL必知必会

## 了解SQL

数据库是一个以某种组织方式存储的数据集合
表是一种结构化的文件，可用来存储某种特定类型的数据
模式（schema） 关于数据库和表的布局及特性的信息

主键 <=> 非空, 唯一
    任意两行都不具有相同的主键值
    每个行都必须具有一个主键值（主键列不允许NULL值）
    不更新主键列中的值
    不用使用可能会更改的列中做主键

## MySQL简介

    DBMS可分为两类：一类为基于共享文件系统的DBMS，另一类为基于客户机—服务器的DBMS

## 使用MySQL

```bash
mysql -u root -p -h localhost -P 3306
USE DATABASENAME;
SHOW DATABSASES;
SHOW TABLES;
SHOW COLUMNS FROM TABLENAME; <=> DESCRIBE TABLENAME;
SHOE STATUS;
SHOW CREATE DATABASE;
SHOW CREATE TABLE;
SHOW GRANTS;
SHOW ERRORS;
SHOW WARNINGS;
HELP SHOW; // 显示允许的SHOW语句
```

## 检索数据

SQL不区分大小写:
    许多SQL开发人员喜欢对所有SQL关键字使用大写，而对所有列和表名使用小写，这样做使代码更易于阅读和调试。

数据表示
    SQL语句一般返回原始的、无格式的数据。数据的格式化是一个表示问题，而不是一个检索问题。因此，表示（对齐和显示上面的价格值，用货币符号和逗号表示其金额）一般在显示该数据的应用程序中规定。一般很少使用实际检索出的原始数据（没有应用程序提供的格式）。
通配符(*)
    一般，除非你确实需要表中的每个列，否则最好别使用*通配符。虽然使用通配符可能会使你自己省事，不用明确列出所需列，但检索不需要的列通常会降低检索和应用程序的性能。
SELECT
    SELECT prod_id FROM products;
    SELECT prod_id, prod_name, prod_price FROM products;
    SELECT * FROM products;
    SELECT DISTINCT vend_id FROM products;

DISTINCT
    不能部分使用DISTINCT,DISTINCT关键字应用于所有列而不仅是前置它的列。如果给出SELECT DISTINCT vend_id,prod_price，除非指定的两个列都不同，否则所有行都将被检索出来。

LIMIT
    下标, 数目
    下表从0开始,即行0
    SELECT prod_name FROM products LIMIT 5;
    SELECT prod_name FROM products LIMIT 1, 1;

    LIMIT 4 OFFSET 3意为从行3开始取4行，就像LIMIT 3, 4一样。

完全限定
    使用库名, 表名限定查询

## 排序检索数据

子句（clause）
    SQL语句由子句构成，有些子句是必需的，而有的是可选的。一个子句通常由一个关键字和所提供的数据组
ORDER BY
    默认升序(ASC),降序(DESC)
    单个:
        SELECT prod_name FROM products ORDER BY prod_price;升序
        SELECT prod_name FROM products ORDER BY prod_name DESC;降序
    多个:
        SELECT prod_name, prod_price, prod_id FROM products ORDER BY prod_price, prod_name;
        首先按价格，然后再按名称排序。
        仅在多个行具有相同的prod_price值时才对产品按prod_name进行排序
        在多个列上降序排序 如果想在多个列上进行降序排序，必须对每个列指定DESC关键字。
    区分大小写和排序顺序 
        在对文本性的数据进行排序时，A与a相同吗？a位于B之前还是位于Z之后？这些问题不是理论问题，其答案取决于数据库如何设置。在字典（dictionary）排序顺序中，A被视为与a相同，这是MySQL（和大多数数据库管理系统）的默认行为。但是，许多数据库管理员能够在需要时改变这种行为（如果你的数据库包含大量外语字符，可能必须这样做）。
    查找最大的价格
        SELECT prod_price FROM products ORDER BY prod_price DESC LIMIT 1;
        在给出ORDER BY子句时，应该保证它位于FROM子句之后。如果使用LIMIT，它必须位于ORDER BY之后。

## 过滤数据

    SQL过滤与应用过滤 数据也可以在应用层过滤。为此目的，SQL的SELECT语句为客户机应用检索出超过实际所需的数据，然后客户机代码对返回数据进行循环，以提取出需要的行。通常，这种实现并不令人满意。因此，对数据库进行了优化，以便快速有效地对数据进行过滤。让客户机应用（或开发语言）处理数据库的工作将会极大地影响应用的性能，并且使所创建的应用完全不具备可伸缩性。此外，如果在客户机上过滤数据，服务器不得不通过网络发送多余的数据，这将导致网络带宽的浪费。

    在同时使用ORDER BY和WHERE子句时，应该让ORDER BY位于WHERE之后

    何时使用引号:
        单引号用来限定字符串。如果将值与串类型的列进行比较，则需要限定引号。用来与数值列进行比较的值不用引号。

    = != <> < <= > >= BETWEEN AND
    IS
    NULL
        在过滤数据时，一定要验证返回数据中确实给出了被过滤列具有NULL的行。

## 数据过滤

    操作符（operator） 用来联结或改变WHERE子句中的子句的关键字。也称为逻辑操作符（logical operator）

    AND, OR
    SELECT prod_name, prod_price FROM products WHERE vend_id = 1003 AND prod_price < 10;
    SELECT prod_name, prod_price FROM products WHERE vend_id = 1002 or vend_id = 1003;

    在WHERE子句中使用圆括号 任何时候使用具有AND和OR操作符的WHERE子句，都应该使用圆括号明确地分组操作符。不要过分依赖默认计算次序，即使它确实是你想要的东西也是如此。使用圆括号没有什么坏处，它能消除歧义。

    IN
        SELECT prod_name, prod_price FROM products WHERE vend_id IN(1002, 1003);
        为什么要使用IN操作符？ 
            1.在使用长的合法选项清单时，IN操作符的语法更清楚且更直观。
            2.在使用IN时，计算的次序更容易管理（因为使用的操作符更少）。
            3.IN操作符一般比OR操作符清单执行更快。
            4.IN的最大优点是可以包含其他SELECT语句，使得能够更动态地建立WHERE子句。第14章将对此进行详细介绍。

    NOT
        为什么使用NOT？
            对于简单的WHERE子句，使用NOT确实没有什么优势。但在更复杂的子句中，NOT非常有用的。例如，在与IN操作符联合使用时，NOT使找出与条件列表不匹配的行非常简单。
        MySQL中的NOT 
            MySQL支持使用NOT 对IN 、BETWEEN 和EXISTS子句取反，这与多数其他DBMS允许使用NOT对各种条件取反有很大的差别。

## 通配符进行过滤

    通配符（wildcard） 用来匹配值的一部分的特殊字符。
    搜索模式（search pattern）由字面值、通配符或两者组合构成的搜索条件。
    通配符本身实际是SQL的WHERE子句中有特殊含义的字符。


    LIKE    
        SELECT prod_id, prod_name FROM products WHERE prod_name LIKE 'jet%';

    通配符
        %, _,
        %并不能匹配NULL

    技巧:
        不要过度使用通配符
        在确实需要使用通配符时，除非绝对有必要，否则不要把它们用在搜索模式的开始处
        仔细注意通配符的位置

## 正则表达式进行搜索

    SELECT prod_name FROM products WHERE prod_name REGEXP 'JetPack .000';
    分区分大小写:
        BINARY
    SELECT prod_name FROM products WHERE prod_name REGEXP BINARY 'JetPack .000';

    | [] \\r(转义字符) [:alnum:](匹配字符类) 元字符 定位符

    测试正则表达式
        SELECT 'hello' REGXP '[0-9]';
        SELECT Trim('abc');
        SELECT 2 * 3;
        SELECT Now();

## 创建计算字段

    Concat()
        拼接（concatenate） 将值联结到一起构成单个值。
        SELECT Concat(vend_name, '(', vend_country, ')') FROM vendors ORDER BY vend_name;

    Trim(), RTrim(), LTrim()
        SELECT Concat(RTrim(vend_name), '(', vend_country, ')') FROM vendors ORDER BY vend_name;

    alias别名
        SELECT Concat(RTrim(vend_name), '(', vend_country, ')') AS vend_title FROM vendors ORDER BY vend_name;
    算术运算
        + - * /
        SELECT prod_id, quantity, item_price FROM orderitems WHERE order_num = 20005;
        SELECT prod_id, quantity, item_price, quantity * item_price AS expanded_price FROM orderitems WHERE order_num = 20005;

##    使用数据处理函数

    Left() 返回串左边的字符
    Length() 返回串的长度
    Locate() 找出串的一个子串
    Lower() 将串转换为小写
    LTrim() 去掉串左边的空格
    Right() 返回串右边的字符
    RTrim() 去掉串右边的空格
    Soundex() 返回串的SOUNDEX值
    SubString() 返回子串的字符
    Upper() 将串转换为大写

    SELECT prod_name, Upper(prod_name) AS prod_name_upcase FROM products;

    SELECT cust_name, cust_contact FROM customers WHERE cust_contact = 'Y. Lie';
    SELECT cust_name, cust_contact FROM customers WHERE SOUNDEX(cust_contact) = SOUNDEX('Y. Lie');


    日期时间
        SELECT cust_id, order_num FROM orders WHERE order_date = '2005-09-01';
        SELECT cust_id, order_num FROM orders WHERE Date(order_date) = '2005-09-01';
        SELECT cust_id, order_num FROM orders WHERE Date(order_date) BETWEEN '2005-09-01' AND '2005-09-30';
        SELECT cust_id, order_num FROM orders WHERE Year(order_date) = 2005 AND Month(order_date) = 9;

    数值函数
        Abs() 返回一个数的绝对值
        Cos() 返回一个角度的余弦
        Exp() 返回一个数的指数值
        Mod() 返回除操作的余数
        Pi() 返回圆周率
        Rand() 返回一个随机数
        Sin() 返回一个角度的正弦
        Sqrt() 返回一个数的平方根
        Tan() 返回一个角度的正切

## 汇总数据

    聚集函数（aggregate function） 运行在行组上，计算和返回单个值的函数。
        AVG() 返回某列的平均值
        COUNT() 返回某列的行数
        MAX() 返回某列的最大值
        MIN() 返回某列的最小值
        SUM() 返回某列值之和
    SELECT AVG(prod_price) AS avg_price FROM products;
    AVG()函数忽略列值为NULL的行。

    SELECT COUNT(*) AS num_cust FROM customers;
    SELECT COUNT(cust_email) AS num_cust FROM customers;
    NULL值 如果指定列名，则指定列的值为空的行被COUNT()函数忽略，但如果COUNT()函数中用的是星号（*），则不忽略。

    SELECT MAX(prod_price) AS max_price FROM products; 
    对非数值数据使用MAX() 虽然MAX()一般用来找出最大的数值或日期值，但MySQL允许将它用来返回任意列中的最大值，包括返回文本列中的最大值。在用于文本数据时，如果数据按相应的列排序，则MAX()返回最后一行。
    NULL值 MAX()函数忽略列值为NULL的行。

    SELECT SUM(quantity) AS item_ordered FROM orderitems WHERE order_num = 20005;
    NULL值 SUM()函数忽略列值为NULL的行。

    DISTINCT
        SELECT AVG(DISTINCT prod_price) AS avg_price FROM products;

## 分组数据

    GROUP BY
        SELECT vend_id, COUNT(*) FROM products GROUP BY vend_id;
        规定:
            1.GROUP BY子句可以包含任意数目的列。这使得能对分组进行嵌套，为数据分组提供更细致的控制。
            2.如果在GROUP BY子句中嵌套了分组，数据将在最后规定的分组上    进行汇总。换句话说，在建立分组时，指定的所有列都一起计算    （所以不能从个别的列取回数据）。
            3.GROUP BY子句中列出的每个列都必须是检索列或有效的表达式    （但不能是聚集函数）。如果在SELECT中使用表达式，则必须在
            4.GROUP BY子句中指定相同的表达式。不能使用别名。 除聚集计算语句外，SELECT语句中的每个列都必须在GROUP BY子句中给出。
            5.如果分组列中具有NULL值，则NULL将作为一个分组返回。如果列    中有多行NULL值，它们将分为一组。
             6.GROUP BY子句必须出现在WHERE子句之后，ORDER BY子句之前。

    HAVING
        目前为止所学过的所有类型的WHERE子句都可以用HAVING来替代。唯一的差别是WHERE过滤行，而HAVING过滤分组。

        WHERE在数据分组前进行过滤，HAVING在数据分组后进行过滤。这是一个重要的区别，WHERE排除的行不包括在分组中。这可能会改变计算值，从而影响HAVING子句中基于这些值过滤掉的分组。

        SELECT cust_id, COUNT(*) FROM orders GROUP BY cust_id HAVING COUNT(*) >= 2;
        SELECT vend_id, COUNT(*) FROM products WHERE prod_price >= 10 GROUP BY vend_id HAVING COUNT(*) >= 2;

    ORDER BY与GROUP BY
        ORDER BY
            排序产生的输出                          
            任意列都可以使用（甚至非选择的列也可以使用）
            不一定需要    
        GROUP BY        
            分组行。但输出可能不是分组的顺序
            只可能使用选择列或表达式列，而且必须使用每个选择列表达式
            如果与聚集函数一起使用列（或表达式），则必须使用

    SELECT子句顺序
        子 句         说 明                         是否必须使用
        SELECT         要返回的列或表达式             是
        FROM         从中检索数据的表             仅在从表选择数据时使用
        WHERE         行级过滤                     否
        GROUP BY     分组说明                     仅在按组计算聚集时使用
        HAVING         组级过滤                     否
        ORDER BY     输出排序顺序                 否
        LIMIT         要检索的行数                 否

## 子查询

    SELECT cust_name, cust_contact FROM customers WHERE cust_id IN (
        SELECT cust_id FROM orders WHERE order_num IN (
            SELECT order_num FROM orderitems WHERE prod_id = 'TNT2'));


    相关子查询（correlated subquery） 涉及外部查询的子查询。
    SELECT cust_name, cust_state, (SELECT COUNT(*) FROM orders WHERE orders.cust_id = customers.cust_id) AS orders FROM customers ORDER BY cust_name;

## 连接表

    WHERE
        SELECT vend_name, prod_name, prod_price FROM vendors, products WHERE vendors.vend_id = products.vend_id ORDER BY vend_name, prod_name;

    笛卡儿积（cartesian product） 由没有联结条件的表关系返回的结果为笛卡儿积。检索出的行的数目将是第一个表中的行数乘以第二个表中的行数。

    内部连接
        SELECT vend_name, prod_name, prod_price FROM vendors INNER JOIN products ON vendors.vend_id = products.vend_id ORDER BY vend_name, prod_name;
    使用哪种语法
        ANSI SQL规范首选INNER JOIN语法。此外，尽管使用WHERE子句定义联结的确比较简单，但是使用明确的联结语法能够确保不会忘记联结条件，有时候这样做也能影响性能。
    多个表连接
        SELECT prod_name, vend_name, prod_price, quantity FROM orderitems, products, vendors WHERE products.vend_id = vendors.vend_id AND orderitems.prod_id = products.prod_id AND order_num = 20005;
    正如第14章所述，子查询并不总是执行复杂SELECT操作的最有效的方法，下面是使用联结的相同查询：
        SELECT cust_name, cust_contact FROM customers, orders, orderitems WHERE customers.cust_id = orders.cust_id AND orders.order_num = orderitems.order_num AND prod_id = 'TNT2';

## 高级连接

    表别名只在查询执行中使用。与列别名不一样，表别名不返回到客户机。

    等值联结
    自然联结<=>内部连接                
    外部联结，左外连接，右外连接，全连接

    SELECT p1.prod_id, p1.prod_name FROM products AS p1, products AS p2 WHERE p1.vend_id = p2.vend_id AND p2.prod_id = 'DTNTR';
        用自联结而不用子查询 自联结通常作为外部语句用来替代从相同表中检索数据时使用的子查询语句。虽然最终的结果是相同的，但有时候处理联结远比处理子查询快得多。应该试一下两种方法，以确定哪一种的性能更好。

    SELECT c.*, o.order_num, o.order_date, oi.prod_id, oi.quantity, oi.item_price FROM customers AS c, orders AS o, orderitems AS oi WHERE c.cust_id = o.cust_id AND oi.order_num = o.order_num AND prod_id = 'FB';
        自然联结排除多次出现，使每个列只返回一次。迄今为止我们建立的每个内部联结都是自然联结，很可能我们永远都不会用到不是自然联结的内部联结。

    SELECT customers.cust_id, orders.order_num FROM customers LEFT OUTER JOIN orders ON customers.cust_id = orders.cust_id;
    SELECT customers.cust_id, orders.order_num FROM customers RIGHT OUTER JOIN orders ON customers.cust_id = orders.cust_id;
        外部联结的类型 存在两种基本的外部联结形式：左外部联结和右外部联结。它们之间的唯一差别是所关联的表的顺序不同。换句话说，左外部联结可通过颠倒FROM或WHERE子句中表的顺序转换为右外部联结。因此，两种类型的外部联结可互换使用，而究竟使用哪一种纯粹是根据方便而定。

## 组合查询

    多数SQL查询都只包含从一个或多个表中返回数据的单条SELECT语句。MySQL也允许执行多个查询（多条SELECT语句），并将结果作为单个查询结果集返回。这些组合查询通常称为并（union）或复合查询（compound query）。

    有两种基本情况，其中需要使用组合查询：
    在单个查询中从不同的表返回类似结构的数据；
    对单个表执行多个查询，按单个查询返回数据。

    如果遵守了这些基本规则或限制，则可以将并用于任何数据检索任务。
        1.UNION必须由两条或两条以上的SELECT语句组成，语句之间用关键字UNION分隔（因此，如果组合4条SELECT语句，将要使用3个UNION关键字）。
        2.UNION中的每个查询必须包含相同的列、表达式或聚集函数（不过各个列不需要以相同的次序列出）。
        3.列数据类型必须兼容：类型不必完全相同，但必须是DBMS可以隐含地转换的类型（例如，不同的数值类型或不同的日期类型）。
    UNION从查询结果集中自动去除了重复的行，这是UNION的默认行为，但是如果需要，可以改变它。事实上，如果想返回所有匹配行，可使用UNION ALL而不是UNION。
    SELECT语句的输出用ORDER BY子句排序。在用UNION组合查询时，只能使用一条ORDER BY子句，它必须出现在最后一条SELECT语句之后。

## 全文本搜索

    两个最常使用的引擎为MyISAM和InnoDB，前者支持全文本搜索，而后者不支持。

    不要在导入数据时使用FULLTEXT
        更新索引要花时间，虽然不是很多，但毕竟要花时间。如果正在导入数据到一个新表，此时不应该启用FULLTEXT索引。应该首先导入所有数据，然后再修改表，定义FULLTEXT。这样有助于更快地导入数据（而且使索引数据的总时间小于在导入每行时分别进行索引所需的总时间）。
    使用两个函数Match()和Against()执行全文本搜索
     SELECT note_text FROM productnotes WHERE MATCH(note_text) Against('rabbit');
    使用完整的Match() 说明 传递给Match() 的值必须与FULLTEXT()定义中的相同。如果指定多个列，则必须列出它们（而且次序正确）。

    搜索不区分大小写除非使用BINARY方式（本章中没有介绍），否则全文本搜索不区分大小写。

    正如所见，全文本搜索提供了简单LIKE搜索不能提供的功能。而且，由于数据是索引的，全文本搜索还相当快。

    查询扩展
    布尔文本搜索

## 插入数据

    总是使用列的列表 一般不要使用没有明确给出列的列表的INSERT语句。使用列的列表能使SQL代码继续发挥作用，即使表结构发生了变化。

    提高INSERT的性能 此技术可以提高数据库处理的性能，因为MySQL用单条INSERT语句处理多个插入比使用多条INSERT语句快。
    INSERT SELECT中SELECT语句可包含WHERE子句以过滤插入的数据。

## 更新和删除数据

    删除表的内容而不是表，DELETE语句从表中删除行，甚至是删除表中所有行。但是，DELETE不删除表本身。
    更快的删除 
        如果想从表中删除所有行，不要使用DELETE。可使用TRUNCATE TABLE语句，它完成相同的工作，但速度更快（TRUNCATE实际是删除原来的表并重新创建一个表，而不是逐行删除表中的数据）

## 创建和操纵表

    InnoDB是一个可靠的事务处理引
    MEMORY在功能等同于MyISAM，但由于数据存储在内存（不是磁盘）中，速度很快（特别适合于临时表）
    MyISAM是一个性能极高的引擎，它支持全文本搜索，但不支持事务处理。

    外键不能跨引擎 混用引擎类型有一个大缺陷。外键（用于强制实施引用完整性，如第1章所述）不能跨引擎，即使用一个引擎的表不能引用具有使用不同引擎的表的外键。

## 视图

    视图是虚拟的表。与包含数据的表不一样，视图只包含使用时动态检索数据的查询。

    视图用CREATE VIEW语句来创建。使用SHOW CREATE VIEW viewname；来查看创建视图的语句。 用DROP删除视图，其语法为DROP VIEW viewname; 更新视图时，可以先用DROP再用CREATE，也可以直接用CREATE OR REPLACE VIEW。如果要更新的视图不存在，则第2条更新语句会创建一个视图；如果要更新的视图存在，则第2条更新语句会替换原有视图。

    本章许多例子中的视图都是不可更新的。这听上去好像是一个严重的限制，但实际上不是，因为视图主要用于数据检索。应该将视图用于检索（SELECT语句）而不用于更新（INSERT、UPDATE和DELETE）。

## 存储过程

    存储过程简单来说，就是为以后的使用而保存的一条或多条MySQL语句的集合。可将其视为批文件，虽然它们的作用不仅限于批处理。
    换句话说，使用存储过程有3个主要的好处，即简单、安全、高性能。

## 游标

    指向查询的结果集

## 触发器

    它们都需要在某个表发生更改时自动处理。这确切地说就是触发器。触发器是MySQL响应以下任意语句而自动执行的一条MySQL语句（或位于BEGIN和END语句之间的一组语句）：
    DELETE；
    INSERT；
    UPDATE。
    其他MySQL语句不支持触发器。

## 事务

    事务处理（transaction processing）可以用来维护数据库的完整性，它保证成批的MySQL操作要么完全执行，要么完全不执行。

    事务（transaction）指一组SQL语句；
         回退（rollback）指撤销指定SQL语句的过程；
        提交（commit）指将未存储的SQL语句结果写入数据库表；
    保留点（savepoint）指事务处理中设置的临时占位符（placeholder），你可以对它发布回退（与回退整个事务处理不同）

    哪些语句可以回退？ 
        事务处理用来管理INSERT、UPDATE和DELETE语句。你不能回退SELECT语句。（这样做也没有什么意义。）你不能回退CREATE或DROP操作。事务处理块中可以使用这两条语句，但如果你执行回退，它们不会被撤销。


    隐含事务关闭当COMMIT或ROLLBACK语句执行后，事务会自动关闭（将来的更改会隐含提交）。
    释放保留点保留点在事务处理完成（执行一条ROLLBACK或COMMIT）后自动释放。自MySQL 5以来，也可以用RELEASE SAVEPOINT明确地释放保留点。

    START TANSACTION    ROLLBACK     COMMIT     SAVEPOINT    ROLLBACK TO

    标志为连接专用 autocommit标志是针对每个连接而不是服务器的。
    SET autocommit = 0

    ACID，是指数据库管理系统（DBMS）在写入或更新资料的过程中，为保证事务（transaction）是正确可靠的，所必须具备的四个特性：原子性（atomicity，或称不可分割性）、一致性（consistency）、隔离性（isolation，又称独立性）、持久性（durability）。

## 全球化和本地化

```sh
SHOW CHARACTER SET; SHOW COLLATION;
SHOW VARIABLES LIKE 'character%';
SHOW VARIABLES LIKE 'collation%';
```

## 安全管理

安全

 ```sh
CREATE USER rrf IDENTIFIED BY 'password';
RENAME USER rrf TO ruifeng;
DROP USER rrf;
SHOW GRANTS FOR rrf;
GRANT SELECT ON test.* TO rrf;
REVOKE SELECT ON test.* TO rrf;
 ```

整个服务器，使用GRANT ALL和REVOKE ALL；
整个数据库，使用ON database.*；
    特定的表，使用ON database.table；
    特定的列；
特定的存储过程。

SET PASSWORD FOR rrf = Password('123456');

## 数据库维护

    FLUSH TABLES;
    ANALYZE TABLE orders;
    日志

## 改善性能

    应该总是使用正确的数据类型。
        决不要检索比需求还要多的数据。换言之，不要用SELECT *（除非你真正需要每个列）。
    有的操作（包括INSERT）支持一个可选的DELAYED关键字，如果使用它，将把控制立即返回给调用程序，并且一旦有可能就实际执行该操作。
    在导入数据时，应该关闭自动提交。你可能还想删除索引（包括FULLTEXT索引），然后在导入完成后再重建它们。
    LIKE很慢。一般来说，最好是使用FULLTEXT而不是LIKE。
    最重要的规则就是，每条规则在某些条件下都会被打破。
