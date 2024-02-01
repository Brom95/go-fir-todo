package main

import (
	"log"

	"github.com/google/uuid"
	"github.com/livefir/fir"
)

func index(storage TaskStorage) fir.RouteFunc {

	return func() fir.RouteOptions {

		return fir.RouteOptions{
			fir.ID("Todo"),
			fir.Content("./templates/index.html"),
			fir.Layout("./templates/layout.html"),
			fir.OnEvent("add", func(ctx fir.RouteContext) error {
				item := &Task{}
				ctx.Bind(item)
				storage.Add(*item)
				log.Println("Task added:", item)

				return ctx.Data(map[string]any{"tasks": storage.GetAll()})
			}),
			fir.OnEvent("complete", func(ctx fir.RouteContext) error {
				sId := new(TaskIdResponse)
				if err := ctx.Bind(sId); err != nil {
					return err
				}
				id, err := uuid.Parse(sId.TaskId)
				if err != nil {
					return err
				}
				err = storage.UpdateStatus(id, true)
				if err != nil {
					return err
				}
				return ctx.Data(map[string]any{"tasks": storage.GetAll()})

			}),
			fir.OnEvent("notcomplete", func(ctx fir.RouteContext) error {
				sId := new(TaskIdResponse)

				if err := ctx.Bind(sId); err != nil {
					return err
				}
				id, err := uuid.Parse(sId.TaskId)
				if err != nil {
					return err
				}
				err = storage.UpdateStatus(id, false)
				if err != nil {
					return err
				}
				return ctx.Data(map[string]any{"tasks": storage.GetAll()})

			}),
			fir.OnLoad(func(ctx fir.RouteContext) error {
				return ctx.Data(map[string]any{"tasks": storage.GetAll()})
			}),
		}

	}
}
