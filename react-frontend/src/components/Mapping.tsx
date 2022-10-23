import { Grid } from '@material-ui/core';

type Props = {};
export const Mapping = (props: Props) => {
   return (
      <Grid container>
         <Grid item xs={12} sm={3}> Formulário </Grid>
         <Grid item xs={12} sm={9}> Mapa </Grid>
      </Grid>
   );
};
