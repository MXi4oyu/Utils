# Set

## 自己封装的Set

## 用法

>
```
func main()  {

   //初始化
   s:=New()

   //添加
   s.Add("宅男帮")
   s.Add(1234567890)
   s.Add("http://www.zhainanbang.net")
   s.Add(true)
   s.Add(3.14)
   //打印集中的元素和长度
   fmt.Println(s.List(),s.Len())
   s.Add(3.14)
   //打印集中的元素和长度
   fmt.Println(s.List(),s.Len())
   //移除
   s.Remove(3.14)
   //打印集中的元素和长度
   fmt.Println(s.List(),s.Len())
   //判断某个元素是否在集合中
   fmt.Println("宅男帮是否在集合中:",s.Has("宅男帮"))
   //移除
   s.Remove("http://www.zhainanbang.net")
   s.Remove(1234567890)
   //判断是否为空
   fmt.Println("集合是否为空:",s.IsEmpty())
   //打印集中的元素和长度
   fmt.Println(s.List(),s.Len())
   //清空
   s.Clear()
   //判断是否为空
   fmt.Println("集合是否为空:",s.IsEmpty())
   //打印集中的元素和长度
   fmt.Println(s.List(),s.Len())



}
```