#question：
我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？


答：应该wrap error 抛给上层。不wrap 这个error, 在上层只知道这个原始错误，对我们去定位问题是不友好的，我们应该携带更多信息给调用者