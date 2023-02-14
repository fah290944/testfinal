package testfinal

import(
	"testing"
	"github.com/asaskevich/govalidator"
	"github.com/onsi/gomega"
	"gorm.io/gorm"
)

type Dara struct{
	gorm.Model
	Name string `valid:"required~h"`
	Email string
}

func TestE(t *testing.T){
	g := gomega.NewGomegaWithT(t)

	t.Run("g", func(t *testing.T){
		e := Dara{
			Name: "",
			Email: "e@gmail.com",
		}
	
		ok,err := govalidator.ValidateStruct(e)
		g.Expect(ok).NotTo(gomega.BeTrue())
		g.Expect(err).NotTo(gomega.BeNil())
		g.Expect(err.Error()).To(gomega.Equal("h"))
	})

}