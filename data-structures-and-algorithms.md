# Data structures and algorithms

## Sorting

### BuiltIn

```go
s := []int{50, 10, 5, 4, 3, 2, 1}
sort.Ints(s)
```

```javascript
let items1 = [50, 10, 5, 4, 3, 2, 1];

items1.sort(function(a, b) {
    return a - b;
});

items1.sort(function(a, b) {
    return b - a;
});

items1.sort((a, b) => a - b);

items1.sort((a, b) => b - a);
```

```javascript
let str = 'zyx';
let array = str.split('');
array.sort((a,b)=>{ return a > b ? 1 : -1});
str = array.join('');
```

```javascript
let items = [[1, 2], [2, 3], [0, 1], [4, 7]];
items.sort(function (a, b) {
    if (a[0] === b[0]) return 0;
    if (a[0] < b[0]) return -1;
    return 1;
});
```

### QuickSort

```go
func QuickSort(items []int) {
	if len(items) > 1 {
		pivot := partition(items)
		QuickSort(items[:pivot])
		QuickSort(items[pivot+1:])
	}
}

func partition(items []int) int {
	end := len(items) - 1
	index, pivot := 0, items[end]
	for i := 0; i < end; i++ {
		if items[i] < pivot {
			items[index], items[i] = items[i], items[index]
			index++
		}
	}
	items[index], items[end] = items[end], items[index]
	return index
}
```

```javascript
let quickSort = function(items) {
    return quick(items, 0, items.length - 1);
}

let quick = function(items, left, right) {
    if(!items) return;
    if(1 < items.length) {
        let p = partition(items, left, right);
        if(left < p - 1) {
            quick(items, left, p-1);
        }
        if(p + 1 < right) {
            quick(items, p+1, right);
        }
    }
    return items;
};

let partition = function(items, left, right) {
    let pivot = items[right];
    let index = left;
    for(let i = left; i < right; i++) {
        if(items[i] < pivot) {
            swap(items, i, index);
            index++;
        }
    }
    swap(items, index, right);
    return index;
};

let swap = function(items, i, j) {
    let t = items[i];
    items[i] = items[j];
    items[j] = t;
};
```

---

## Search

### BuiltIn

```go
// ascending order
func Search(items []int, target int) bool {
	i := sort.Search(len(items), func(i int) bool { return items[i] >= target })
	if i < len(items) && items[i] == target {
		return true
	}
	return false
}
```

### BinarySearch

```go
func BinarySearch(items []int, target int) int {
	l, r := 0, len(items)-1
	for l <= r {
		mid := (l + r) / 2
		if items[mid] == target {
			return mid
		}
		if items[mid] < target {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return -1
}
```

---

## Array

### go's array
```go

```

### javascript's array

```javascript
let array1 = [];
let array2 = new Array();

for (index in array1) {

}

for (element of array1) {

}

// array.splice(index, numberOfDelete, data1, data2, ...)

```

---

## Tree

### Traversal

#### InOrder

```go
// recursive

// stack + iterative
```

#### PreOrder
```go
// recursive

// stack + iterative
```

#### PostOrder
```go
// recursive

// stack + iterative + append to result's first position + avoid append nil
```

### LevelOrder
```go
// queue + iterative + iterative

```

### Search

### Insert

```go
// recursive

// iterative

```

### delete

```go
// iterative + find successor

```

### find min node

### find max node

### maximum depth of tree

```go
// recursive

// iterative

```

### minimum depth of tree

```go
// recursive

// iterative + return count when the leaf node is encountered

```

## String

### go's string

#### substr
- substr := str[i:j] 
- substr := str[i:] 
- substr := str[:j] 
- substr := string([]rune(str)[i:j])
- substr := string([]rune(str)[i:])
- substr := string([]rune(str)[:j])

#### slice
- str[i:j]
- str[i:i+l]

#### []byte or []rune
- []byte(str)
- []rune(str)

#### index
- strings.Index(s, substr)

#### trim
- strings.TrimSpace()

#### join
- strings.Join(slice, "")

#### type casting 
- number, err := strconv.Atoi(str)
- str := strconv.Itoa(number)

### case converter
- strings.ToLower(str)
- strings.ToUpper(str)

#### builder
```go
// var b strings.Builder
// b.Reset()
// b.Write()
// b.WriteByte()
// b.WriteRune()
// b.WriteString()
```

### javascript's string

#### substr
- str.substr(i, len)

#### slice
- str.slice(i, j)

#### trim
- str.trim()

#### join
strArray.join('')

#### type casting 
- number = Number("101")
- number = parseInt("101")
- number = parseFloat("101.101")

## Dynamic programming

- 能否達成
- 多少種方法
- 極大值極小值
