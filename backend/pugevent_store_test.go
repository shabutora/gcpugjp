package backend

import (
	"fmt"
	"testing"
	"time"

	"github.com/favclip/testerator"
)

func TestPugEventStore_Upsert(t *testing.T) {
	_, ctx, err := testerator.SpinUp() // gae/pythonのインスタンスが無ければ起動、あれば使いまわす
	if err != nil {
		t.Fatal(err.Error())
	}
	defer testerator.SpinDown() // プロセスをシャットダウンせずに、Datastoreなどの内容をクリアする

	store, err := NewPugEventStore(ctx)
	if err != nil {
		t.Fatalf("%+v", err)
	}

	now := time.Now()
	candidates := []struct {
		param PugEvent
	}{
		{
			param: PugEvent{
				OrganizationID: "tokyo",
				Title:          "GCPUG Day",
				Description:    "GCPUGやるぞー！",
				URL:            "https://gcpug.jp",
				Limit:          10,
				Accepted:       10,
				StartAt:        now,
				EndAt:          now,
			},
		},
		{
			param: PugEvent{
				OrganizationID: "tokyo",
				Title:          "GCPUG Day",
				Description:    "GCPUGやるぞー！",
				URL:            "https://gcpug.jp",
				Limit:          20,
				Accepted:       10,
				StartAt:        now,
				EndAt:          now,
			},
		},
	}

	for _, v := range candidates {
		e := v.param

		stored, err := store.Upsert(ctx, &e)
		if err != nil {
			t.Fatalf("%+v", err)
		}
		if stored.Key == nil {
			t.Fatalf("Key is Empty.")
		}
		if e, g := e.OrganizationID, stored.OrganizationID; e != g {
			t.Fatalf("expected OrganizationID %s; got %s", e, g)
		}
		if e, g := e.Title, stored.Title; e != g {
			t.Fatalf("expected Title %s; got %s", e, g)
		}
		if e, g := e.Description, stored.Description; e != g {
			t.Fatalf("expected Description %s; got %s", e, g)
		}
		if e, g := e.URL, stored.URL; e != g {
			t.Fatalf("expected URL %s; got %s", e, g)
		}
		if e, g := e.Limit, stored.Limit; e != g {
			t.Fatalf("expected Limit %d; got %d", e, g)
		}
		if e, g := e.Accepted, stored.Accepted; e != g {
			t.Fatalf("expected Accepted %d; got %d", e, g)
		}
		//if e, g := e.StartAt, stored.StartAt; !EqualTime(e, g) {
		//	t.Fatalf("expected StartAt %s; got %s", e, g)
		//}
		//if e, g := e.EndAt, stored.EndAt; !EqualTime(e, g) {
		//	t.Fatalf("expected EndAt %s; got %s", e, g)
		//}
		if stored.CreatedAt.IsZero() {
			t.Fatalf("CreatedAt is Zero")
		}
		if stored.UpdatedAt.IsZero() {
			t.Fatalf("UpdatedAt is Zero")
		}
		if stored.SchemaVersion == 0 {
			t.Fatalf("SchemaVersion is Zero")
		}
	}
}

func TestPugEventStore_Get(t *testing.T) {
	_, ctx, err := testerator.SpinUp() // gae/pythonのインスタンスが無ければ起動、あれば使いまわす
	if err != nil {
		t.Fatal(err.Error())
	}
	defer testerator.SpinDown() // プロセスをシャットダウンせずに、Datastoreなどの内容をクリアする

	store, err := NewPugEventStore(ctx)
	if err != nil {
		t.Fatalf("%+v", err)
	}

	e := PugEvent{
		OrganizationID: "tokyo",
		Title:          "GCPUG Day",
		Description:    "GCPUGやるぞー！",
		URL:            "https://gcpug.jp",
		StartAt:        time.Now(),
		EndAt:          time.Now(),
	}

	pe, err := store.Upsert(ctx, &e)
	if err != nil {
		t.Fatalf("%+v", err)
	}

	stored, err := store.Get(ctx, pe.Key)
	if err != nil {
		t.Fatalf("%+v", err)
	}
	if stored.Key == nil {
		t.Fatalf("Key is Empty.")
	}
	if e, g := e.OrganizationID, stored.OrganizationID; e != g {
		t.Fatalf("expected OrganizationID %s; got %s", e, g)
	}
	if e, g := e.Title, stored.Title; e != g {
		t.Fatalf("expected Title %s; got %s", e, g)
	}
	if e, g := e.Description, stored.Description; e != g {
		t.Fatalf("expected Description %s; got %s", e, g)
	}
	if e, g := e.URL, stored.URL; e != g {
		t.Fatalf("expected URL %s; got %s", e, g)
	}
	//if e, g := e.StartAt, stored.StartAt; !EqualTime(e, g) {
	//	t.Fatalf("expected StartAt %s; got %s", e, g)
	//}
	//if e, g := e.EndAt, stored.EndAt; !EqualTime(e, g) {
	//	t.Fatalf("expected EndAt %s; got %s", e, g)
	//}
	if stored.CreatedAt.IsZero() {
		t.Fatalf("CreatedAt is Zero")
	}
	if stored.UpdatedAt.IsZero() {
		t.Fatalf("UpdatedAt is Zero")
	}
	if stored.SchemaVersion == 0 {
		t.Fatalf("SchemaVersion is Zero")
	}
}

func TestPugEventStore_List(t *testing.T) {
	_, ctx, err := testerator.SpinUp() // gae/pythonのインスタンスが無ければ起動、あれば使いまわす
	if err != nil {
		t.Fatal(err.Error())
	}
	defer testerator.SpinDown() // プロセスをシャットダウンせずに、Datastoreなどの内容をクリアする

	store, err := NewPugEventStore(ctx)
	if err != nil {
		t.Fatalf("%+v", err)
	}

	candidates := []struct {
		limit       int
		length      int
		existCursor bool
	}{
		{1, 1, true},
		{1, 1, true},
		{2, 2, true},
	}

	for i, v := range candidates {
		e := PugEvent{
			OrganizationID: "tokyo",
			Title:          "GCPUG Day",
			Description:    "GCPUGやるぞー！",
			URL:            fmt.Sprintf("https://gcpug.jp/%d", i),
			StartAt:        time.Now(),
			EndAt:          time.Now(),
		}

		_, err = store.Upsert(ctx, &e)
		if err != nil {
			t.Fatalf("%+v", err)
		}

		param := &PugEventListParam{
			Limit: v.limit,
		}
		results, err := store.List(ctx, param)
		if err != nil {
			t.Fatalf("i = %d, failed PugEventStore.List %+v", i, err)
		}
		if e, g := v.length, len(results.List); e != g {
			t.Fatalf("i = %d, expected results.List %d; got %d", i, e, g)
		}
		existCursor := results.NextCursor != nil
		if e, g := v.existCursor, existCursor; e != g {
			t.Fatalf("i = %d, expected results.NextCursor %t; got %t", i, e, g)
		}
		if results.List[0].KeyStr == "" {
			t.Fatalf("i = %d, expected results.List[0].KeyStr is not empty", i)
		}
	}
}
