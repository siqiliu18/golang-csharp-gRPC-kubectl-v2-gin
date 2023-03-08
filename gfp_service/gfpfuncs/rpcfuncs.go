package gfpfuncs

import (
	"context"
	"fmt"
	"net/http"

	templates "gfpservice/merge-proto"

	"github.com/gin-gonic/gin"
)

// RetrieveTemplate is the same
func (s *Service) RetrieveTemplate(ctx context.Context, in *templates.TemplateRequest) (*templates.TemplateResponse, error) {
	fmt.Println(">>> RetrieveTemplate func")

	req := templates.TemplateRequest {
		Templatename: in.GetTemplatename(),
		Templatetype: in.GetTemplatetype(),
		Cbedna: in.GetCbedna(),
	}

	res, err := s.MergeClient.RetrieveTemplate(ctx, &req)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	out := templates.TemplateResponse {
		Documenturi: res.GetDocumenturi(),
	}

	return &out, nil
}

// GuiFunc is gateway
func (s *Service) GuiFunc(ctx *gin.Context) {
	// _, err := strconv.ParseInt(ctx.Param("a"), 10, 64)
	// if err != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameter A"})
	// 	return
	// }

	// _, err = strconv.ParseInt(ctx.Param("b"), 10, 64)
	// if err != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{"error": "Invalid parameter B"})
	// 	return
	// }

	in := templates.TemplateRequest {
		Templatename: ctx.Param("a"),
		Templatetype: ctx.Param("b"),
		Cbedna: "G123",
	}

	if res, err := s.RetrieveTemplate(ctx, &in); err == nil {
		ctx.JSON(http.StatusOK, gin.H{
			"result": fmt.Sprint(res.GetDocumenturi()),
		})
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
}