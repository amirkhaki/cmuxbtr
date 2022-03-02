# cmuxbtr
amazon scraper


## نحوه استفاده 
#### افزودن
برای اینکه یک محصول در وبسایتتون قیمتش بر اساس قیمت یک محصول در آمازون، آپدیت بشه،
، از دستور زیر استفاده کنین 
```
.\cmuxbtr.exe connect --id <ID> -u "<URL>" -t <T> -w <W>
```
استفاده کنین، بجای ` <ID> ` شناسه محصول توی ووکامرس رو قرار بدین،
  بجای ` <URL> ` لینک محصول توی Amazon.ae رو قرار بدین 
  بجای ` <T> ` درصدی که باید به قیمت محصول اضافه بشه رو قرار بدین 
  بجای ` <W> ` کارمزدتون رو وارد کنین 
  مثلا: 
  ```
 .\cmuxbtr.exs connect --id 123 -u "https://www.amazon.ae/dp/B08TSKZSH9/" -t 3 -w 300000 
  ```
  
  
#### حذف کردن 
از دستور زیر برای حذف کردن یک محصول از دیتابیس برنامه استفاده کنین، در نتیجه دیگه آپدیت نخواهد شد 
```
.\cmuxbtr.exe delete --id <ID>
```
که `<ID>` شناسه محصول هست.
#### لیست محصولات 
با دستور زیر میتونین لیست محصولات رو ببینین 
```
.\cmuxbtr.exe list 
``` 

#### آپدیت کردن محصولات 
با استفاده از دستور زیر یا اجرای مستقیم برنامه میتونین قبمت ها رو آپدیت کنین توی وبسایت 
```
.\cmuxbtr.exe update 
```
اگر میخواین که فقط یک محصول مشخص آپدیت بشه، از دستور زیر استفاده کنین 
```
.\cmuxbtr.exe update --id <ID> 
``` 
که `<ID>` شناسه محصول هست .
