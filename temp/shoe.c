#include <stdio.h>
#include <string.h>
#include <stdlib.h>
#include <assert.h>

// 枚举
typedef enum _ShowEnum
{
    LI_NING,    // 鞋子种类 - 李宁
    HUI_LI,     // 鞋子种类 - 回力
    QIAO_DAN,   // 鞋子种类 - 乔丹
    YANG_JICHENG,    // 鞋子种类 - new
} ShowEnum;

// 商店结构
typedef struct _Shoe
{
    ShowEnum type;  // 鞋子种类申明
    int price;      // 参数 - 鞋子价格
    char *strtype;  // 鞋子种类指针
    void (*Shoe_Type_StrCopy)(struct _Shoe *);  // 函数申明
    void (*SetPrice)(struct _Shoe *, int num);  // 函数申明
    void (*Printf_Shoe_Show)(struct _Shoe *);   // 函数申明
    void (*Init)(struct _Shoe *);               // 初始化函数申明
} Shoe;

void shoe_type_strCopy(Shoe *_this)
{
    if (_this->type == LI_NING)
    {
        _this->strtype = "LI_NING";
    }
    else if (_this->type == HUI_LI)
    {
        _this->strtype = "HUI_LI";
    }
    else if (_this->type == QIAO_DAN)
    {
        _this->strtype = "QIAO_DAN";
    }
}

void setprice(Shoe *_this, int num)
{
    _this->price = num;
}

void printf_shoe_show(Shoe *_this)
{
    printf("\nShoe name:%s,price:%d\n", _this->strtype, _this->price);
}

void Shoe_Init(Shoe *_this)
{
    _this->Shoe_Type_StrCopy = shoe_type_strCopy;
    _this->SetPrice = setprice;
    _this->Printf_Shoe_Show = printf_shoe_show;
}

int main()
{
    // Shoe *s = (Shoe *)malloc(sizeof(Shoe));

    Shoe s;
    s.Init = Shoe_Init;
    s.Init(&s);

    s.type = LI_NING;
    s.Shoe_Type_StrCopy(&s);
    s.SetPrice(&s, 80);
    s.Printf_Shoe_Show(&s);

    s.type = QIAO_DAN;
    s.Shoe_Type_StrCopy(&s);
    s.SetPrice(&s, 130);
    s.Printf_Shoe_Show(&s);

    printf("\n");
    return 0;
}