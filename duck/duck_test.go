package duck

import (
	"fmt"
	"testing"

	gomock "github.com/golang/mock/gomock"
)

func TestBreed(t *testing.T) {
	ctrl := gomock.NewController()
	defer ctrl.Finish()

	duck := NewMockDuckI(ctrl)
	quack := "quack quack"
	duck.EXPECT().quack().Return(quack)

	farmer := NewFarmer(duck)
	fmt.Println(farmer)
	breed := farmer.Breed()
	if fmt.Sprintf("#0 %s\n", quack) != breed {
		t.Error("breeding is failed [" + breed + "]")
	}

}
