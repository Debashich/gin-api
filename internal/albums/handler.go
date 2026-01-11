package albums

import (
    "net/http"

    "github.com/gin-gonic/gin"
)

type Handler struct {
    store *Store
}

func NewHandler(store *Store) *Handler {
    return &Handler{store: store}
}

// GET /albums
func (h *Handler) GetAll(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, h.store.GetAll())
}

// GET /albums/:id
func (h *Handler) GetByID(c *gin.Context) {
    id := c.Param("id")

    album, found := h.store.GetByID(id)
    if !found {
        c.JSON(http.StatusNotFound, gin.H{"error": "album not found"})
        return
    }

    c.IndentedJSON(http.StatusOK, album)
}

// POST /albums
func (h *Handler) Add(c *gin.Context) {
    var newAlbum Album

    if err := c.BindJSON(&newAlbum); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid request"})
        return
    }

    h.store.Add(newAlbum)
    c.IndentedJSON(http.StatusCreated, newAlbum)
}
