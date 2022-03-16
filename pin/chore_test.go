package pin

import (
	"github.com/stretchr/testify/require"
	"math/big"
	"testing"
)

//
////TODO: size based calculation
//func Test_calculateUntil(t *testing.T) {
//	dateFormat := "2006-01-02 15:04:05"
//	from, err := time.Parse(dateFormat, "2022-03-15 03:00:00")
//	require.NoError(t, err)
//
//	//price: 2E5 / day / byte
//	// 2E5 * 2.5E5 = 5E11 per day
//	// paid: 1E13 / 5E11 -->  100 / 5 days --> 20
//	require.Equal(t, "2022-04-04 03:00:00", calculateUntil(from, big.NewInt(200000), big.NewInt(10_000_000_000_000), 2_500_000).Format(dateFormat))
//}

func Test_calculateUntil(t *testing.T) {
	price, _ := new(big.Int).SetString("2731", 10) // 0.01 / year

	require.Equal(t, uint(5*366)+1, calculateDays(price, big.NewInt(5_000_000), 2_500_000))
}
