package component

import (
	"testing"
	"time"

	"github.com/shiwano/submarine/server/battle/server/battle/context"
	"github.com/shiwano/submarine/server/battle/server/resource"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/ungerik/go3d/float64/vec2"
)

const timeLayout = "15:04:05.000"

func TestMotor(t *testing.T) {
	Convey("Motor", t, func() {
		stageMesh, _ := resource.Loader.LoadMesh(1)
		lightMap, _ := resource.Loader.LoadLightMap(1)
		c := context.NewContext(stageMesh, lightMap)
		c.Now, _ = time.Parse(timeLayout, "00:00:00.000")
		m := NewMotor(c, &vec2.T{1, 1}, 0, 3, 5*time.Second)

		Convey("#ToAPIType", func() {
			Convey("should returns a Movement message", func() {
				message := m.ToAPIType(1)
				So(message.ActorId, ShouldEqual, 1)
				So(message.Coerce(), ShouldBeNil)
			})
		})

		Convey("#Position", func() {
			Convey("when the accelerator is shutdown", func() {
				Convey("should return the initial position", func() {
					So(m.Position(), ShouldResemble, &vec2.T{1, 1})
				})

				Convey("1 second later", func() {
					c.Now, _ = time.Parse(timeLayout, "00:00:01.000")

					Convey("should return the initial position", func() {
						So(m.Position(), ShouldResemble, &vec2.T{1, 1})
					})
				})

				Convey("5 seconds later", func() {
					c.Now, _ = time.Parse(timeLayout, "00:00:05.000")

					Convey("should return the initial position", func() {
						So(m.Position(), ShouldResemble, &vec2.T{1, 1})
					})
				})

				Convey("10 seconds later", func() {
					c.Now, _ = time.Parse(timeLayout, "00:00:10.000")

					Convey("should return the initial position", func() {
						So(m.Position(), ShouldResemble, &vec2.T{1, 1})
					})
				})
			})

			Convey("when the accelerator is accelerating", func() {
				m.Accelerate(m.Position())

				Convey("should return the initial position", func() {
					So(m.Position(), ShouldResemble, &vec2.T{1, 1})
				})

				Convey("1 second later", func() {
					c.Now, _ = time.Parse(timeLayout, "00:00:01.000")

					Convey("should return the calculated position", func() {
						So(m.Position()[0], ShouldAlmostEqual, 1+0.3)
						So(m.Position()[1], ShouldAlmostEqual, 1)
					})
				})

				Convey("5 seconds later", func() {
					c.Now, _ = time.Parse(timeLayout, "00:00:05.000")

					Convey("should return the calculated position", func() {
						So(m.Position()[0], ShouldAlmostEqual, 1+7.5)
						So(m.Position()[1], ShouldAlmostEqual, 1)
					})
				})

				Convey("10 seconds later", func() {
					c.Now, _ = time.Parse(timeLayout, "00:00:10.000")

					Convey("should return the calculated position", func() {
						So(m.Position()[0], ShouldAlmostEqual, 1+7.5+3*5)
						So(m.Position()[1], ShouldAlmostEqual, 1)
					})
				})
			})

			Convey("when the accelerator is stopping", func() {
				m.Accelerate(m.Position())
				c.Now, _ = time.Parse(timeLayout, "00:00:05.000")
				m.Brake(m.Position())

				Convey("should return the initial position", func() {
					So(m.Position(), ShouldResemble, &vec2.T{8.5, 1})
				})

				Convey("1 second later", func() {
					c.Now, _ = time.Parse(timeLayout, "00:00:06.000")

					Convey("should return the calculated position", func() {
						So(m.Position()[0], ShouldAlmostEqual, 8.5+2.7)
						So(m.Position()[1], ShouldAlmostEqual, 1)
					})
				})

				Convey("5 seconds later", func() {
					c.Now, _ = time.Parse(timeLayout, "00:00:10.000")

					Convey("should return the calculated position", func() {
						So(m.Position()[0], ShouldAlmostEqual, 8.5+7.5)
						So(m.Position()[1], ShouldAlmostEqual, 1)
					})
				})

				Convey("10 seconds later", func() {
					c.Now, _ = time.Parse(timeLayout, "00:00:15.000")

					Convey("should return the calculated position", func() {
						So(m.Position()[0], ShouldAlmostEqual, 8.5+7.5)
						So(m.Position()[1], ShouldAlmostEqual, 1)
					})
				})
			})

			Convey("when the accelerator is accelerating from the middle", func() {
				m.Accelerate(m.Position())
				c.Now, _ = time.Parse(timeLayout, "00:00:03.000")
				m.Brake(m.Position())
				c.Now, _ = time.Parse(timeLayout, "00:00:05.000")
				m.Accelerate(m.Position())

				Convey("should return the initial position", func() {
					So(m.Position()[0], ShouldAlmostEqual, 6.1)
					So(m.Position()[1], ShouldAlmostEqual, 1)
				})

				Convey("4 seconds later", func() {
					c.Now, _ = time.Parse(timeLayout, "00:00:09.000")

					Convey("should return the calculated position", func() {
						So(m.Position()[0], ShouldAlmostEqual, 13.3)
						So(m.Position()[1], ShouldAlmostEqual, 1)
					})
				})

				Convey("9 seconds later", func() {
					c.Now, _ = time.Parse(timeLayout, "00:00:14.000")

					Convey("should return the calculated position", func() {
						So(m.Position()[0], ShouldAlmostEqual, 13.3+3*5)
						So(m.Position()[1], ShouldAlmostEqual, 1)
					})
				})
			})

			Convey("when the accelerator is stopping from the middle", func() {
				m.Accelerate(m.Position())
				c.Now, _ = time.Parse(timeLayout, "00:00:03.000")
				m.Brake(m.Position())

				Convey("should return the initial position", func() {
					So(m.Position()[0], ShouldAlmostEqual, 3.7)
					So(m.Position()[1], ShouldAlmostEqual, 1)
				})

				Convey("3 seconds later", func() {
					c.Now, _ = time.Parse(timeLayout, "00:00:06.000")

					Convey("should return the calculated position", func() {
						So(m.Position()[0], ShouldAlmostEqual, 6.4)
						So(m.Position()[1], ShouldAlmostEqual, 1)
					})
				})

				Convey("8 seconds later", func() {
					c.Now, _ = time.Parse(timeLayout, "00:00:11.000")

					Convey("should return the calculated position", func() {
						So(m.Position()[0], ShouldAlmostEqual, 6.4)
						So(m.Position()[1], ShouldAlmostEqual, 1)
					})
				})
			})

			Convey("when the direction changed", func() {
				m.Turn(m.Position(), 90)
				m.Accelerate(m.Position())
				c.Now, _ = time.Parse(timeLayout, "00:00:01.000")

				Convey("should return the calculated position", func() {
					So(m.Position()[0], ShouldAlmostEqual, 1)
					So(m.Position()[1], ShouldAlmostEqual, 1.3)
				})

				Convey("and the direction changed once more", func() {
					m.Turn(m.Position(), 0)
					c.Now, _ = time.Parse(timeLayout, "00:00:06.000")

					Convey("should return the calculated position", func() {
						So(m.Position()[0], ShouldAlmostEqual, 11.2)
						So(m.Position()[1], ShouldAlmostEqual, 1.3)
					})
				})
			})

			Convey("when the accelerator is idling", func() {
				m.Accelerate(m.Position())
				c.Now, _ = time.Parse(timeLayout, "00:00:01.000")
				m.Idle(m.Position())

				Convey("should return the idling position", func() {
					So(m.Position()[0], ShouldAlmostEqual, 1+0.3)
					So(m.Position()[1], ShouldAlmostEqual, 1)
				})

				Convey("some seconds later", func() {
					c.Now, _ = time.Parse(timeLayout, "00:00:05.000")

					Convey("should return the idling position", func() {
						So(m.Position()[0], ShouldAlmostEqual, 1+0.3)
						So(m.Position()[1], ShouldAlmostEqual, 1)
					})
				})

				Convey("when the motor stop idling", func() {
					c.Now, _ = time.Parse(timeLayout, "00:00:05.000")
					m.Accelerate(m.Position())
					c.Now, _ = time.Parse(timeLayout, "00:00:06.000")

					Convey("should keep accelerating", func() {
						So(m.Position()[0], ShouldAlmostEqual, 1+0.3+3)
						So(m.Position()[1], ShouldAlmostEqual, 1)
					})
				})
			})
		})
	})
}
