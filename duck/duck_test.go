package duck

import (
	"fmt"
	"testing"

	gomock "github.com/golang/mock/gomock"
)

func TestBreed(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockDuck := NewMockDuckI(ctrl)
	quack := "quack quack test"
	mockDuck.EXPECT().quack().Return(quack)

	mockFarmer := NewFarmer(mockDuck)
	fmt.Println(mockFarmer)
	breed := mockFarmer.Breed()
	if fmt.Sprintf("#0 %s\n", quack) != breed {
		t.Error("breeding is failed [" + breed + "]")
	}

}
