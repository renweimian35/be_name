> 小孩出生，名字伴随一生，头等大事，岂能儿戏？遂写一程序，成事在人，谋事在随机吧。
### 词汇源整理
- 楚辞：1
- 诗经：2
- 唐诗：3
- 宋词：4
- 论语：5
- 曹操：6

### 初步需求如下
- 通过以上词汇源，输出数字id，表示从中获取名字，可支持多个。
- 通过输出姓名字数，来随机匹配文字
- 显示名字的来源，出处。
初步需求效果如：
```shell
E:\workspaces\goland\be-name> go run main.go -count=1 -num=2
子来   取自:
南有樛木,甘瓠累之.君子有酒,嘉宾式燕绥之.
翩翩者鵻,烝然来思.君子有酒,嘉宾式燕又思.
```
### 迭代需求一
先补充背景知识
#### 名字的技巧
- 声韵：就是为了使名字叫起来在声韵方面好听、响亮。最应该注意一点，即慎重取双声叠韵的名字，双声即两个字的声母相同，叠韵即两个字的韵母相同。
  如韩凡楠（三字的韵母都是an），念起来就不是那么顺畅，有点拗口的感觉。
- 声调：即平仄，一二声为平，三四声为仄，平仄相间才会读起来悦耳。所以，姓名中也一样，尽量做到汉字的平仄相间，**结尾字宜以平声结尾**，因为平声调感觉响亮。
- 不要使用生僻字
- 不要使用多音字，让人读起来无所适从，遇到不必要的麻烦
- 名字笔画不要太多，娃把名字抄写一百遍的时候，会诅咒你
- 可添加虚字（而、何、乎、乃、从 其、若、所、为、也、以、因、于、与、则 者、之、可、如、亦、尔 、科、在、再、客、然、以 、已、大、不、自、准
  才、仅、今、非、允、又、有、由、必、竟、予、充、少、宜、更、复）等等，有种朦胧感。
#### 需求
- 支持录入姓氏，通过姓氏的声韵和声调，确定首位名字，如阳姓，则首字排除an、ang音韵字符。如复姓，则取姓氏的最后一字作为基准。
- 确定随机到的名字为平仄相间。如阳读第二声为平，则第二个字为仄，第三字为平，如：阳(yáng)顶(dǐng)天(tiān)
- 笔画多余15的，不要。
#### 效果如下：
```shell
C:\Users\44312\AppData\Local\JetBrains\GoLand2023.2\tmp\GoLand\___8go_build_main_go.exe
阳小存(yángxiǎocún),出自 孔子 的 论语-泰伯篇 ：
曾子有疾，召门弟子曰：“启予足，启予手。《诗》云：‘战战兢兢，如临深渊，如履薄冰。’而今而后，吾知免夫，小子！”
曾子有疾，孟敬子问之。曾子言曰：“鸟之将死，其鸣也哀；人之将死，其言也善。君子所贵乎道者三：动容貌，斯远暴慢矣；正颜色，斯近信矣；出辞气，斯远鄙倍矣。笾豆之事，则有司存。”
---------------------------------------------------------------
阳利人(yánglìrén),出自 孔子 的 论语-子罕篇 ：
子罕言利与命与仁。
达巷党人曰：“大哉孔子！博学而无所成名。”子闻之，谓门弟子曰：“吾何执？执御乎，执射乎？吾执御矣。”
---------------------------------------------------------------
阳勉如(yángmiǎnrú),出自 孔子 的 论语-子罕篇 ：
子曰：“出则事公卿，入则事父兄，丧事不敢不勉，不为酒困，何有于我哉？”
子在川上曰：“逝者如斯夫！不舍昼夜。”
```
### 需求迭代二
通过上述可以得出基础名字，但感觉又不太像是名字，可用性不强，再次增加需求
- 增加汉字解释，通过释义来倒推名字，如表示缓慢、行走等等词语，推断出 徐行 二字，来自苏东坡的定风波，何妨吟啸且徐行。
- 增加名的词组搜索，如果有解释，则名字也比较通顺。
如上，应该可以大大提高名字可用性。
- 增加看过的名字，就不要再次循环到黑名单。
#### 效果如下:
```shell
 阳此心(yángcǐxīn),出自 杜甫 的 唐诗-登楼 ：
花近高楼伤客心，万方多难此登临。
其中 此心 的含义为：《人同此心，心同此理》 指合情合理的事，大家想法都会相同。   
---------------------------------------------------------------
 阳力争(yánglìzhēng),出自 曹操 的 曹操诗集-蒿里行 ：
军合力不齐，踌躇而雁行。
势利使人争，嗣还自相戕。
其中 力争 的含义为：《力争上游》 上游河的上流，比喻先进的地位。努力奋斗，争取先进再先进。 
---------------------------------------------------------------
 阳志闻(yángzhìwén),出自 孔子 的 论语-季氏篇 ：
孔子曰：“见善如不及，见不善如探汤；吾见其人矣。吾闻其语矣。隐居以求其志，行义以达其道；吾闻其语矣，未见其人也。”
其中 志闻 的含义为：《博闻强志》 形容知识丰富，记忆力强。
---------------------------------------------------------------
 阳径幽(yángjìngyōu),出自 常建 的 唐诗-题破山寺后禅院 ：
曲径通幽处，禅房花木深。
其中 径幽 的含义为：《曲径通幽》 曲弯曲；径小路；幽指深远僻静之处。弯曲的小路，通到幽深僻静的地方。
```
剩下的，就是自己肉眼去看了。
#### 感谢：
https://github.com/chinese-poetry/chinese-poetry 提供词汇源  

https://github.com/mozillazg/go-pinyin 提供拼音、声韵、声调

https://github.com/pwxcoo/chinese-xinhua 提供汉字释义/笔画、成语大全



