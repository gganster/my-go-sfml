# Démarrage rapide
## 1.Win - Environnement windows
- télécharger CSFML2.5.1, placer le dossier décompressé à la racine (C:/CSFML_2.5.1). Le chemin d'accès doit être exact.
- installer mingw. La commande gcc doit être accessible dans le terminal (redémarrer la console si nécessaire)
- passer à l'étape 2.Win

## 1.Unix Environnement mac os
//TODO

## 2.Win Compilation
```sh
./build.bat
./sfml.exe
```
passer à l'étape 3 pour la documentation

## 2.Unix
//TODO

## 3.Documentation

### Boucle de rendu
La boucle de rendu est néccessaire pour gérer des fenêtres
Elle se compose toujours de 5 étapes:
- pollEvent (My_Window_update)
- traitement des évènements et mise à jour de l'état
- effacemement de la fenêtre
- dessin des éléments sur l'écran
- affichage de l'écran

1 tour de cette boucle est appelé une frame. D'où le nom de FPS (Frame per second)

### Drawable
Vous avez accès a 2 élements à dessiner:
- MyRect
- MyCircle

Ils ont chacuns leur propre propriété (taille, position, couleur, ...) Et doivent tous les 2 avoir leur fonction _display appelé pour être affiché à l'écran

#### MyRect
```go
func MyRect_create() *MyRect
func MyRect_destroy(o *MyRect)
func MyRect_display(o *MyRect, w *MyWindow)

func MyRect_setX(o *MyRect, x float32)
func MyRect_setY(o *MyRect, y float32) 
func MyRect_setWidth(o *MyRect, w float32)
func MyRect_setHeight(o *MyRect, h float32)
func MyRect_setColor(o *MyRect, r, g, b byte)
func MyRect_setPosition(o *MyRect, x, y float32)

func MyRect_getWidth(o *MyRect) float32
func MyRect_getHeight(o *MyRect) float32
func MyRect_getPosition(o *MyRect) (float32, float32)
func MyRect_getColor(o *MyRect) (byte, byte, byte)
func MyRect_getX(o *MyRect) float32
func MyRect_getY(o *MyRect) float32 
```

#### MyCircle
```go
func MyCircle_create() *MyCircle
func MyCircle_destroy(o *MyCircle)
func MyCircle_display(o *MyCircle, w *MyWindow)

func MyCircle_setX(o *MyCircle, x float32)
func MyCircle_setY(o *MyCircle, y float32) 
func MyCircle_setRadius(o *MyCircle, r float32)
func MyCircle_setColor(o *MyCircle, r, g, b byte)
func MyCircle_setPosition(o *MyCircle, x, y float32)

func MyCircle_getRadius(o *MyCircle) float32
func MyCircle_getPosition(o *MyCircle) (float32, float32)
func MyCircle_getColor(o *MyCircle) (byte, byte, byte)
func MyCircle_getX(o *MyCircle) float32
func MyCircle_getY(o *MyCircle) float32
```

### Evenements
Les évènements sont disponibles dans l'object Window via les fonctions d'accès.

A la différence des autres, les fonction *Once attendront le relachement de la touche avant de renvoyer un autre true

```go
//sourie
func MyWindow_getMouseX(o *MyWindow) float32
func MyWindow_getMouseY(o *MyWindow) float32
func MyWindow_isMouseLeftPressed(o *MyWindow) bool
func MyWindow_isMouseRightPressed(o *MyWindow) bool
func MyWindow_isMouseLeftPressedOnce(o *MyWindow) bool
func MyWindow_isMouseRightPressedOnce(o *MyWindow) bool

//clavier
func MyWindow_isArrowUpPressed(o *MyWindow) bool
func MyWindow_isArrowDownPressed(o *MyWindow) bool
func MyWindow_isArrowLeftPressed(o *MyWindow) bool
func MyWindow_isArrowRightPressed(o *MyWindow) bool
func MyWindow_isSpacePressed(o *MyWindow) bool
func MyWindow_isAPressed(o *MyWindow) bool
func MyWindow_isBPressed(o *MyWindow) bool
func MyWindow_isCPressed(o *MyWindow) bool
... A-Z

func MyWindow_isArrowUpPressedOnce(o *MyWindow) bool
... Vous avez compris le concept
```

Un code de test vous est fournis dans main.go