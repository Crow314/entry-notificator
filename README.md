# 寺子屋deおんらいん

## API

### Child

#### POST `/children`
Register child

```json
{
  "first_name": "太郎",
  "family_name": "山田",
  "birth_date": "2009-09-10",
  "address": {
    "postal_code": "639-1080",
    "prefecture_code": 29,
    "city": "大和郡山市",
    "street": "矢田町22番地",
    "room": ""
  },
  "contact": {
    "email": "yamada@example.com",
    "phone_number": "090-0000-0000"
  }
}
```

See [全国地方公共団体コード#都道府県コード](https://ja.wikipedia.org/wiki/%E5%85%A8%E5%9B%BD%E5%9C%B0%E6%96%B9%E5%85%AC%E5%85%B1%E5%9B%A3%E4%BD%93%E3%82%B3%E3%83%BC%E3%83%89#%E9%83%BD%E9%81%93%E5%BA%9C%E7%9C%8C%E3%82%B3%E3%83%BC%E3%83%89) for `prefecture_code`

### Card

#### POST `/cards`

Register id card for the child

```json
{
  "child_id": "aae61ee1-40a3-44c4-829a-49c68e227082",
  "token": "0101020c2927e5df",
  "token_type": "FeliCa"
}
```

#### PUT `/card/:id/scan`

Change entry status of the child

### Supporter

#### POST `/supporters`
Register supporter

```json
{
  "first_name": "留太",
  "family_name": "高専",
  "birth_date": "1989-01-08",
  "email": "kosen@example.com",
  "address": {
    "postal_code": "193-0834",
    "prefecture_code": 13,
    "city": "八王子市",
    "street": "東浅川町701-2",
    "room": ""
  }
}
```

See [全国地方公共団体コード#都道府県コード](https://ja.wikipedia.org/wiki/%E5%85%A8%E5%9B%BD%E5%9C%B0%E6%96%B9%E5%85%AC%E5%85%B1%E5%9B%A3%E4%BD%93%E3%82%B3%E3%83%BC%E3%83%89#%E9%83%BD%E9%81%93%E5%BA%9C%E7%9C%8C%E3%82%B3%E3%83%BC%E3%83%89) for `prefecture_code`
