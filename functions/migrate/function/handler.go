package function

import (
	"net/http"

	log "github.com/my-ermes-labs/log"
)

func Handle(w http.ResponseWriter, r *http.Request) {
	log.MyNodeLog(node.AreaName, "\n\nMIGRATE!\n\n")
}
