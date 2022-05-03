package relpaginator_test

import (
	"context"
	"time"

	"github.com/go-rel/rel"
	"github.com/go-rel/rel/where"
	"github.com/go-rel/reltest"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"github.com/manicar2093/relpaginator"
)

type User struct {
	UserUUID  string
	CreatedAt time.Time
}

func (c User) Table() string {
	return "users"
}

var _ = Describe("RelPaginator", func() {

	var (
		repo      *reltest.Repository
		paginable relpaginator.Paginable
		config    relpaginator.RelPaginatorConfig
		dbTable   string

		pageRequested      = uint(1)
		previousPageExpect = uint(0)
		nextPageExpected   = uint(2)
		totalPagesExpected = uint(100)
	)

	BeforeEach(func() {
		repo = reltest.New()
		config = relpaginator.RelPaginatorConfig{
			PageSize: 10,
		}
		paginable = relpaginator.NewPaginator(repo, &config)
		dbTable = "users"
	})

	It("should create a Pagination with its filters info", func() {
		pageSort := relpaginator.PageSort{
			Page:     pageRequested,
			SortBy:   []string{"field_one", "field_two"},
			SortDesc: []bool{true, false},
		}
		userUUIDFilter := where.Eq("user_uuid", "an_uuid")
		userCreatedAtFilter := where.Eq("created_at", time.Now())
		pageSort.SetFiltersQueries(userUUIDFilter, userCreatedAtFilter)
		var users []User
		expected_user_find_all := []User{{}, {}, {}, {}, {}, {}, {}}
		repo.ExpectCount(dbTable, userUUIDFilter, userCreatedAtFilter).Result(1000)
		repo.ExpectFindAll(
			userUUIDFilter,
			userCreatedAtFilter,
			rel.SortDesc("field_one"),
			rel.SortAsc("field_two"),
			rel.Limit(config.PageSize),
			rel.Offset(
				0,
			),
		).Result(expected_user_find_all)

		page, err := paginable.CreatePagination(context.Background(), dbTable, &users, &pageSort)

		Expect(err).ToNot(HaveOccurred())
		Expect(page.TotalEntries).ToNot(BeZero())
		Expect(page.Data).To(Equal(&expected_user_find_all))
		Expect(page.TotalPages).To(Equal(totalPagesExpected))
		Expect(page.NextPage).To(Equal(nextPageExpected))
		Expect(page.PreviousPage).To(Equal(previousPageExpect))
		Expect(page.CurrentPage).To(Equal(pageRequested))
	})

	It("should create a Pagination without filters info", func() {
		pageSort := relpaginator.PageSort{
			Page: pageRequested,
		}
		var users []User
		expected_user_find_all := []User{{}, {}, {}, {}, {}, {}, {}}

		repo.ExpectCount(dbTable).Result(1000)
		repo.ExpectFindAll(
			rel.Limit(config.PageSize),
			rel.Offset(
				0,
			),
		).Result(expected_user_find_all)

		page, err := paginable.CreatePagination(context.Background(), dbTable, &users, &pageSort)

		Expect(err).ToNot(HaveOccurred())
		Expect(page.TotalEntries).ToNot(BeZero())
		Expect(page.Data).To(Equal(&expected_user_find_all))
		Expect(page.TotalPages).To(Equal(totalPagesExpected))
		Expect(page.NextPage).To(Equal(nextPageExpected))
		Expect(page.PreviousPage).To(Equal(previousPageExpect))
		Expect(page.CurrentPage).To(Equal(pageRequested))
	})

	Context("if the last page is requested", func() {
		It("should return the correct data", func() {
			pageRequested = uint(100)
			previousPageExpect = uint(99)
			nextPageExpected = uint(1)
			totalPagesExpected = uint(100)
			pageSort := relpaginator.PageSort{
				Page: pageRequested,
			}
			var users []User
			expected_user_find_all := []User{{}, {}, {}, {}, {}, {}, {}}

			repo.ExpectCount(dbTable).Result(1000)
			repo.ExpectFindAll(
				rel.Limit(config.PageSize),
				rel.Offset(
					990,
				),
			).Result(expected_user_find_all)

			page, err := paginable.CreatePagination(context.Background(), dbTable, &users, &pageSort)

			Expect(err).ToNot(HaveOccurred())
			Expect(page.TotalEntries).ToNot(BeZero())
			Expect(page.Data).To(Equal(&expected_user_find_all))
			Expect(page.TotalPages).To(Equal(totalPagesExpected))
			Expect(page.NextPage).To(Equal(nextPageExpected))
			Expect(page.PreviousPage).To(Equal(previousPageExpect))
			Expect(page.CurrentPage).To(Equal(pageRequested))
		})
	})

	Context("if requested page does not exist", func() {
		It("should return a PageError", func() {
			pageRequested = uint(101)
			pageSort := relpaginator.PageSort{
				Page: pageRequested,
			}
			var users []User
			expected_user_find_all := []User{{}, {}, {}, {}, {}, {}, {}}
			repo.ExpectCount(dbTable).Result(1000)
			repo.ExpectFindAll(
				rel.Limit(config.PageSize),
				rel.Offset(
					990,
				),
			).Result(expected_user_find_all)

			page, err := paginable.CreatePagination(context.Background(), dbTable, &users, &pageSort)

			Expect(page).To(BeNil())
			Expect(err).To(HaveOccurred())
			Expect(err).To(BeAssignableToTypeOf(relpaginator.PageError{}))

		})
	})

	Context("if found entries are less than page size", func() {
		It("should return one page info", func() {
			pageRequested = uint(1)
			previousPageExpect = uint(0)
			nextPageExpected = uint(1)
			totalPagesExpected = uint(1)
			pageSort := relpaginator.PageSort{
				Page: pageRequested,
			}
			var users []User
			expected_user_find_all := []User{{}, {}, {}, {}, {}, {}, {}}
			repo.ExpectCount(dbTable).Result(2)
			repo.ExpectFindAll(
				rel.Limit(config.PageSize),
				rel.Offset(
					0,
				),
			).Result(expected_user_find_all)

			page, err := paginable.CreatePagination(context.Background(), dbTable, &users, &pageSort)

			Expect(err).ToNot(HaveOccurred())
			Expect(page.TotalEntries).ToNot(BeZero())
			Expect(page.Data).To(Equal(&expected_user_find_all))
			Expect(page.TotalPages).To(Equal(totalPagesExpected))
			Expect(page.NextPage).To(Equal(nextPageExpected))
			Expect(page.PreviousPage).To(Equal(previousPageExpect))
			Expect(page.CurrentPage).To(Equal(pageRequested))
		})
	})

})
