package routes
import(fmt)

type Router struct {
	config *Config
	router *chi
}

func NewRouter() *Router{
	return &Router{
		config: NewConfig().SetTimeout(serviceConfig().Timeout),
		router: chi.NewRouter(),
	}

}

func (r *Router) SetRouter() *chi.Mux{

}

func (r *Router)setConfigRouter(){

}

func RouterHealth(){

}

func RouterProduct(){

}

func EnableTimeout(){

}

func EnableCORS(){

}

func EnableRecover(){

}

func EnableRequestID(){

}

func EnableRealIP(){
	
}