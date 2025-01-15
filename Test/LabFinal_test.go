package unit

import (
	"testing" 

	. "github.com/onsi/gomega"
	"github.com/asaskevich/govalidator"
	"github.com/JanisataMJ/FinalLabExam/entity"
)

//done

func TestValid(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run("All Valid", func(t *testing.T) {
		exam := entity.LabFinal{
			SID:	"B6515683",
			Name:	"Janisata",
			No:		21,
		}

		ok,err := govalidator.ValidateStruct(exam)

		g.Expect(ok).To(BeTrue())
		g.Expect(err).To(BeNil())
	})
}

func TestNotRequired(t *testing.T) {
	g := NewGomegaWithT(t)

	t.Run("SID is not required", func(t *testing.T) {
		exam := entity.LabFinal{
			SID:	"",
			Name:	"Janisata",
			No:		21,
		}

		ok,err := govalidator.ValidateStruct(exam)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("SID is not required"))
	})

	t.Run("Name is not required", func(t *testing.T) {
		exam := entity.LabFinal{
			SID:	"65156833",
			Name:	"",
			No:		21,
		}

		ok,err := govalidator.ValidateStruct(exam)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("Name is not required"))
	})

	t.Run("No is not required", func(t *testing.T) {
		exam := entity.LabFinal{
			SID:	"65156833",
			Name:	"Janisata",
			No:		0,
		}

		ok,err := govalidator.ValidateStruct(exam)

		g.Expect(ok).NotTo(BeTrue())
		g.Expect(err).NotTo(BeNil())

		g.Expect(err.Error()).To(Equal("No is not required"))
	})
}